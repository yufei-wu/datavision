package api

import (
	"fmt"

	"github.com/grafana/grafana/pkg/api/dtos"
	"github.com/grafana/grafana/pkg/bus"
	m "github.com/grafana/grafana/pkg/models"
	"github.com/grafana/grafana/pkg/services/alerting"
	"github.com/grafana/grafana/pkg/services/guardian"
)

func ValidateOrgAlert(c *m.ReqContext) {
	id := c.ParamsInt64(":alertId")
	query := m.GetAlertByIdQuery{Id: id}

	if err := bus.Dispatch(&query); err != nil {
		c.JsonApiErr(404, "警告没找到", nil)
		return
	}

	if c.OrgId != query.Result.OrgId {
		c.JsonApiErr(403, "您没有权限编辑或查看警告", nil)
		return
	}
}

func GetAlertStatesForDashboard(c *m.ReqContext) Response {
	dashboardID := c.QueryInt64("dashboardId")

	if dashboardID == 0 {
		return Error(400, "缺少查询参数dashboardId", nil)
	}

	query := m.GetAlertStatesForDashboardQuery{
		OrgId:       c.OrgId,
		DashboardId: c.QueryInt64("dashboardId"),
	}

	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "获取警告状态失败", err)
	}

	return JSON(200, query.Result)
}

// GET /api/alerts
func GetAlerts(c *m.ReqContext) Response {
	query := m.GetAlertsQuery{
		OrgId:       c.OrgId,
		DashboardId: c.QueryInt64("dashboardId"),
		PanelId:     c.QueryInt64("panelId"),
		Limit:       c.QueryInt64("limit"),
		User:        c.SignedInUser,
	}

	states := c.QueryStrings("state")
	if len(states) > 0 {
		query.State = states
	}

	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "列出警告失败", err)
	}

	for _, alert := range query.Result {
		alert.Url = m.GetDashboardUrl(alert.DashboardUid, alert.DashboardSlug)
	}

	return JSON(200, query.Result)
}

// POST /api/alerts/test
func AlertTest(c *m.ReqContext, dto dtos.AlertTestCommand) Response {
	if _, idErr := dto.Dashboard.Get("id").Int64(); idErr != nil {
		return Error(400, "仪表盘在验证警告规则时至少保存一次", nil)
	}

	backendCmd := alerting.AlertTestCommand{
		OrgId:     c.OrgId,
		Dashboard: dto.Dashboard,
		PanelId:   dto.PanelId,
	}

	if err := bus.Dispatch(&backendCmd); err != nil {
		if validationErr, ok := err.(alerting.ValidationError); ok {
			return Error(422, validationErr.Error(), nil)
		}
		return Error(500, "测试规则失败", err)
	}

	res := backendCmd.Result
	dtoRes := &dtos.AlertTestResult{
		Firing:         res.Firing,
		ConditionEvals: res.ConditionEvals,
		State:          res.Rule.State,
	}

	if res.Error != nil {
		dtoRes.Error = res.Error.Error()
	}

	for _, log := range res.Logs {
		dtoRes.Logs = append(dtoRes.Logs, &dtos.AlertTestResultLog{Message: log.Message, Data: log.Data})
	}
	for _, match := range res.EvalMatches {
		dtoRes.EvalMatches = append(dtoRes.EvalMatches, &dtos.EvalMatch{Metric: match.Metric, Value: match.Value})
	}

	dtoRes.TimeMs = fmt.Sprintf("%1.3fms", res.GetDurationMs())

	return JSON(200, dtoRes)
}

// GET /api/alerts/:id
func GetAlert(c *m.ReqContext) Response {
	id := c.ParamsInt64(":alertId")
	query := m.GetAlertByIdQuery{Id: id}

	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "列出警告失败", err)
	}

	return JSON(200, &query.Result)
}

func GetAlertNotifiers(c *m.ReqContext) Response {
	return JSON(200, alerting.GetNotifiers())
}

func GetAlertNotifications(c *m.ReqContext) Response {
	query := &m.GetAllAlertNotificationsQuery{OrgId: c.OrgId}

	if err := bus.Dispatch(query); err != nil {
		return Error(500, "获取警告通知失败", err)
	}

	result := make([]*dtos.AlertNotification, 0)

	for _, notification := range query.Result {
		result = append(result, &dtos.AlertNotification{
			Id:        notification.Id,
			Name:      notification.Name,
			Type:      notification.Type,
			IsDefault: notification.IsDefault,
			Created:   notification.Created,
			Updated:   notification.Updated,
		})
	}

	return JSON(200, result)
}

func GetAlertNotificationByID(c *m.ReqContext) Response {
	query := &m.GetAlertNotificationsQuery{
		OrgId: c.OrgId,
		Id:    c.ParamsInt64("notificationId"),
	}

	if err := bus.Dispatch(query); err != nil {
		return Error(500, "获取警告通知失败", err)
	}

	return JSON(200, query.Result)
}

func CreateAlertNotification(c *m.ReqContext, cmd m.CreateAlertNotificationCommand) Response {
	cmd.OrgId = c.OrgId

	if err := bus.Dispatch(&cmd); err != nil {
		return Error(500, "创建警告通知失败", err)
	}

	return JSON(200, cmd.Result)
}

func UpdateAlertNotification(c *m.ReqContext, cmd m.UpdateAlertNotificationCommand) Response {
	cmd.OrgId = c.OrgId

	if err := bus.Dispatch(&cmd); err != nil {
		return Error(500, "更新警告通知失败", err)
	}

	return JSON(200, cmd.Result)
}

func DeleteAlertNotification(c *m.ReqContext) Response {
	cmd := m.DeleteAlertNotificationCommand{
		OrgId: c.OrgId,
		Id:    c.ParamsInt64("notificationId"),
	}

	if err := bus.Dispatch(&cmd); err != nil {
		return Error(500, "删除警告通知失败", err)
	}

	return Success("通知已删除")
}

//POST /api/alert-notifications/test
func NotificationTest(c *m.ReqContext, dto dtos.NotificationTestCommand) Response {
	cmd := &alerting.NotificationTestCommand{
		Name:     dto.Name,
		Type:     dto.Type,
		Settings: dto.Settings,
	}

	if err := bus.Dispatch(cmd); err != nil {
		if err == m.ErrSmtpNotEnabled {
			return Error(412, err.Error(), err)
		}
		return Error(500, "发送通知失败", err)
	}

	return Success("测试通知发送")
}

//POST /api/alerts/:alertId/pause
func PauseAlert(c *m.ReqContext, dto dtos.PauseAlertCommand) Response {
	alertID := c.ParamsInt64("alertId")

	query := m.GetAlertByIdQuery{Id: alertID}

	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "获取警告失败", err)
	}

	guardian := guardian.New(query.Result.DashboardId, c.OrgId, c.SignedInUser)
	if canEdit, err := guardian.CanEdit(); err != nil || !canEdit {
		if err != nil {
			return Error(500, "Error while checking permissions for Alert", err)
		}

		return Error(403, "Access denied to this dashboard and alert", nil)
	}

	cmd := m.PauseAlertCommand{
		OrgId:    c.OrgId,
		AlertIds: []int64{alertID},
		Paused:   dto.Paused,
	}

	if err := bus.Dispatch(&cmd); err != nil {
		return Error(500, "", err)
	}

	var response m.AlertStateType = m.AlertStatePending
	pausedState := "un-paused"
	if cmd.Paused {
		response = m.AlertStatePaused
		pausedState = "paused"
	}

	result := map[string]interface{}{
		"alertId": alertID,
		"state":   response,
		"message": "Alert " + pausedState,
	}

	return JSON(200, result)
}

//POST /api/admin/pause-all-alerts
func PauseAllAlerts(c *m.ReqContext, dto dtos.PauseAllAlertsCommand) Response {
	updateCmd := m.PauseAllAlertCommand{
		Paused: dto.Paused,
	}

	if err := bus.Dispatch(&updateCmd); err != nil {
		return Error(500, "Failed to pause alerts", err)
	}

	var response m.AlertStateType = m.AlertStatePending
	pausedState := "un paused"
	if updateCmd.Paused {
		response = m.AlertStatePaused
		pausedState = "paused"
	}

	result := map[string]interface{}{
		"state":          response,
		"message":        "alerts " + pausedState,
		"alertsAffected": updateCmd.ResultCount,
	}

	return JSON(200, result)
}
