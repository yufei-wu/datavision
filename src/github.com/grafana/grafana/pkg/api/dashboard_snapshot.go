package api

import (
	"time"

	"github.com/grafana/grafana/pkg/api/dtos"
	"github.com/grafana/grafana/pkg/bus"
	"github.com/grafana/grafana/pkg/metrics"
	m "github.com/grafana/grafana/pkg/models"
	"github.com/grafana/grafana/pkg/services/guardian"
	"github.com/grafana/grafana/pkg/setting"
	"github.com/grafana/grafana/pkg/util"
)

func GetSharingOptions(c *m.ReqContext) {
	c.JSON(200, util.DynMap{
		"externalSnapshotURL":  setting.ExternalSnapshotUrl,
		"externalSnapshotName": setting.ExternalSnapshotName,
		"externalEnabled":      setting.ExternalEnabled,
	})
}

func CreateDashboardSnapshot(c *m.ReqContext, cmd m.CreateDashboardSnapshotCommand) {
	if cmd.Name == "" {
		cmd.Name = "Unnamed snapshot"
	}

	if cmd.External {
		// external snapshot ref requires key and delete key
		if cmd.Key == "" || cmd.DeleteKey == "" {
			c.JsonApiErr(400, "Missing key and delete key for external snapshot", nil)
			return
		}

		cmd.OrgId = -1
		cmd.UserId = -1
		metrics.M_Api_Dashboard_Snapshot_External.Inc()
	} else {
		cmd.Key = util.GetRandomString(32)
		cmd.DeleteKey = util.GetRandomString(32)
		cmd.OrgId = c.OrgId
		cmd.UserId = c.UserId
		metrics.M_Api_Dashboard_Snapshot_Create.Inc()
	}

	if err := bus.Dispatch(&cmd); err != nil {
		c.JsonApiErr(500, "创建快照失败", err)
		return
	}

	c.JSON(200, util.DynMap{
		"key":       cmd.Key,
		"deleteKey": cmd.DeleteKey,
		"url":       setting.ToAbsUrl("dashboard/snapshot/" + cmd.Key),
		"deleteUrl": setting.ToAbsUrl("api/snapshots-delete/" + cmd.DeleteKey),
	})
}

// GET /api/snapshots/:key
func GetDashboardSnapshot(c *m.ReqContext) {
	key := c.Params(":key")
	query := &m.GetDashboardSnapshotQuery{Key: key}

	err := bus.Dispatch(query)
	if err != nil {
		c.JsonApiErr(500, "获取仪表盘快照失败", err)
		return
	}

	snapshot := query.Result

	// expired snapshots should also be removed from db
	if snapshot.Expires.Before(time.Now()) {
		c.JsonApiErr(404, "仪表盘快照未找到", err)
		return
	}

	dto := dtos.DashboardFullWithMeta{
		Dashboard: snapshot.Dashboard,
		Meta: dtos.DashboardMeta{
			Type:       m.DashTypeSnapshot,
			IsSnapshot: true,
			Created:    snapshot.Created,
			Expires:    snapshot.Expires,
		},
	}

	metrics.M_Api_Dashboard_Snapshot_Get.Inc()

	c.Resp.Header().Set("Cache-Control", "public, max-age=3600")
	c.JSON(200, dto)
}

// GET /api/snapshots-delete/:key
func DeleteDashboardSnapshot(c *m.ReqContext) Response {
	key := c.Params(":key")

	query := &m.GetDashboardSnapshotQuery{DeleteKey: key}

	err := bus.Dispatch(query)
	if err != nil {
		return Error(500, "获取仪表盘快照失败", err)
	}

	if query.Result == nil {
		return Error(404, "获取仪表盘快照失败", nil)
	}
	dashboard := query.Result.Dashboard
	dashboardID := dashboard.Get("id").MustInt64()

	guardian := guardian.New(dashboardID, c.OrgId, c.SignedInUser)
	canEdit, err := guardian.CanEdit()
	if err != nil {
		return Error(500, "检查快照权限的过程中出错", err)
	}

	if !canEdit && query.Result.UserId != c.SignedInUser.UserId {
		return Error(403, "拒绝访问该快照", nil)
	}

	cmd := &m.DeleteDashboardSnapshotCommand{DeleteKey: key}

	if err := bus.Dispatch(cmd); err != nil {
		return Error(500, "删除仪表盘快照失败", err)
	}

	return JSON(200, util.DynMap{"message": "快照已删除. CDN 缓存中的快照数据1个小时后会被删除."})
}

// GET /api/dashboard/snapshots
func SearchDashboardSnapshots(c *m.ReqContext) Response {
	query := c.Query("query")
	limit := c.QueryInt("limit")

	if limit == 0 {
		limit = 1000
	}

	searchQuery := m.GetDashboardSnapshotsQuery{
		Name:         query,
		Limit:        limit,
		OrgId:        c.OrgId,
		SignedInUser: c.SignedInUser,
	}

	err := bus.Dispatch(&searchQuery)
	if err != nil {
		return Error(500, "检索失败", err)
	}

	dtos := make([]*m.DashboardSnapshotDTO, len(searchQuery.Result))
	for i, snapshot := range searchQuery.Result {
		dtos[i] = &m.DashboardSnapshotDTO{
			Id:          snapshot.Id,
			Name:        snapshot.Name,
			Key:         snapshot.Key,
			DeleteKey:   snapshot.DeleteKey,
			OrgId:       snapshot.OrgId,
			UserId:      snapshot.UserId,
			External:    snapshot.External,
			ExternalUrl: snapshot.ExternalUrl,
			Expires:     snapshot.Expires,
			Created:     snapshot.Created,
			Updated:     snapshot.Updated,
		}
	}

	return JSON(200, dtos)
}
