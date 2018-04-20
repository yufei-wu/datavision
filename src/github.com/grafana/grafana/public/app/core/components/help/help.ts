import coreModule from '../../core_module';
import appEvents from 'app/core/app_events';

export class HelpCtrl {
  tabIndex: any;
  shortcuts: any;

  /** @ngInject */
  constructor() {
    this.tabIndex = 0;
    this.shortcuts = {
      '全局': [
        { keys: ['g', 'h'], description: '跳转到首页仪表盘' },
        { keys: ['g', 'p'], description: '跳转到资料' },
        { keys: ['s', 'o'], description: '打开搜索' },
        { keys: ['s', 's'], description: '在偏好仪表盘中进行搜索' },
        { keys: ['s', 't'], description: '在标签中进行搜索' },
        { keys: ['esc'], description: '退出编辑/设置页面' },
      ],
      '仪表盘': [
        { keys: ['mod+s'], description: '保存仪表盘' },
        { keys: ['d', 'r'], description: '刷新所有面板' },
        { keys: ['d', 's'], description: '仪表盘设置' },
        { keys: ['d', 'v'], description: '切换处于活动状态/查看模式' },
        { keys: ['d', 'k'], description: '切换信息亭模式（隐藏顶部导航）' },
        { keys: ['d', 'E'], description: '展开所有行' },
        { keys: ['d', 'C'], description: '关闭所有行' },
        { keys: ['mod+o'], description: '切换共享图十字线' },
      ],
      '聚焦的面板': [
        { keys: ['e'], description: '切换面板编辑视图' },
        { keys: ['v'], description: '切换面板全屏视图' },
        { keys: ['p', 's'], description: '打开面板共享模式' },
        { keys: ['p', 'd'], description: '复制面板' },
        { keys: ['p', 'r'], description: '删除面板' },
      ],
      '时间范围': [
        { keys: ['t', 'z'], description: '缩小时间范围' },
        {
          keys: ['t', '<i class="fa fa-long-arrow-left"></i>'],
          description: '将时间范围回移',
        },
        {
          keys: ['t', '<i class="fa fa-long-arrow-right"></i>'],
          description: '将时间范围前推',
        },
      ],
    };
  }

  dismiss() {
    appEvents.emit('hide-modal');
  }
}

export function helpModal() {
  return {
    restrict: 'E',
    templateUrl: 'public/app/core/components/help/help.html',
    controller: HelpCtrl,
    bindToController: true,
    transclude: true,
    controllerAs: 'ctrl',
    scope: {},
  };
}

coreModule.directive('helpModal', helpModal);
