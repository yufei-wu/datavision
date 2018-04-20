import { types, getEnv, flow } from 'mobx-state-tree';
import { ServerStat } from './ServerStat';

export const ServerStatsStore = types
  .model('ServerStatsStore', {
    stats: types.array(ServerStat),
    error: types.optional(types.string, ''),
  })
  .actions(self => ({
    load: flow(function* load() {
      const backendSrv = getEnv(self).backendSrv;
      const res = yield backendSrv.get('/api/admin/stats');
      self.stats.clear();
      self.stats.push(ServerStat.create({ name: '仪表盘总数', value: res.dashboards }));
      self.stats.push(ServerStat.create({ name: '用户总数', value: res.users }));
      self.stats.push(ServerStat.create({ name: '最近30天登录的用户总数', value: res.activeUsers }));
      self.stats.push(ServerStat.create({ name: '组织的总数', value: res.orgs }));
      self.stats.push(ServerStat.create({ name: '播放列表总数', value: res.playlists }));
      self.stats.push(ServerStat.create({ name: '快照总数', value: res.snapshots }));
      self.stats.push(ServerStat.create({ name: '仪表盘标签总数', value: res.tags }));
      self.stats.push(ServerStat.create({ name: '星标仪表盘总数', value: res.stars }));
      self.stats.push(ServerStat.create({ name: '警告总数', value: res.alerts }));
    }),
  }));
