import React from 'react';
import { inject, observer } from 'mobx-react';
import { toJS } from 'mobx';
import PageHeader from 'app/core/components/PageHeader/PageHeader';
import IContainerProps from 'app/containers/IContainerProps';
import { getSnapshot } from 'mobx-state-tree';
import appEvents from 'app/core/app_events';

@inject('nav', 'folder', 'view')
@observer
export class FolderSettings extends React.Component<IContainerProps, any> {
  formSnapshot: any;

  constructor(props) {
    super(props);
    this.loadStore();
  }

  loadStore() {
    const { nav, folder, view } = this.props;

    return folder.load(view.routeParams.get('uid') as string).then(res => {
      this.formSnapshot = getSnapshot(folder);
      view.updatePathAndQuery(`${res.url}/settings`, {}, {});

      return nav.initFolderNav(toJS(folder.folder), 'manage-folder-settings');
    });
  }

  onTitleChange(evt) {
    this.props.folder.setTitle(this.getFormSnapshot().folder.title, evt.target.value);
  }

  getFormSnapshot() {
    if (!this.formSnapshot) {
      this.formSnapshot = getSnapshot(this.props.folder);
    }

    return this.formSnapshot;
  }

  save(evt) {
    if (evt) {
      evt.stopPropagation();
      evt.preventDefault();
    }

    const { nav, folder, view } = this.props;

    folder
      .saveFolder({ overwrite: false })
      .then(newUrl => {
        view.updatePathAndQuery(newUrl, {}, {});

        appEvents.emit('dashboard-saved');
        appEvents.emit('alert-success', ['Folder saved']);
      })
      .then(() => {
        return nav.initFolderNav(toJS(folder.folder), 'manage-folder-settings');
      })
      .catch(this.handleSaveFolderError.bind(this));
  }

  delete(evt) {
    if (evt) {
      evt.stopPropagation();
      evt.preventDefault();
    }

    const { folder, view } = this.props;
    const title = folder.folder.title;

    appEvents.emit('confirm-modal', {
      title: '删除',
      text: `您想要删除该文件夹以及它下面的所有面板吗？`,
      icon: 'fa-trash',
      yesText: '删除',
      onConfirm: () => {
        return folder.deleteFolder().then(() => {
          appEvents.emit('alert-success', ['文件夹删除', `${title} 已经被删除`]);
          view.updatePathAndQuery('dashboards', '', '');
        });
      },
    });
  }

  handleSaveFolderError(err) {
    if (err.data && err.data.status === 'version-mismatch') {
      err.isHandled = true;

      const { nav, folder, view } = this.props;

      appEvents.emit('confirm-modal', {
        title: '冲突',
        text: '有人已经更新了该文件夹.',
        text2: '您仍然想要保存该文件夹吗？',
        yesText: '保存 & 复写',
        icon: 'fa-warning',
        onConfirm: () => {
          folder
            .saveFolder({ overwrite: true })
            .then(newUrl => {
              view.updatePathAndQuery(newUrl, {}, {});

              appEvents.emit('dashboard-saved');
              appEvents.emit('alert-success', ['文件夹已保存']);
            })
            .then(() => {
              return nav.initFolderNav(toJS(folder.folder), 'manage-folder-settings');
            });
        },
      });
    }
  }

  render() {
    const { nav, folder } = this.props;

    if (!folder.folder || !nav.main) {
      return <h2>加载中</h2>;
    }

    return (
      <div>
        <PageHeader model={nav as any} />
        <div className="page-container page-body">
          <h2 className="page-sub-heading">文件夹设置</h2>

          <div className="section gf-form-group">
            <form name="folderSettingsForm" onSubmit={this.save.bind(this)}>
              <div className="gf-form">
                <label className="gf-form-label width-7">名称</label>
                <input
                  type="text"
                  className="gf-form-input width-30"
                  value={folder.folder.title}
                  onChange={this.onTitleChange.bind(this)}
                />
              </div>
              <div className="gf-form-button-row">
                <button
                  type="submit"
                  className="btn btn-success"
                  disabled={!folder.folder.canSave || !folder.folder.hasChanged}
                >
                  <i className="fa fa-save" /> 保存
                </button>
                <button className="btn btn-danger" onClick={this.delete.bind(this)} disabled={!folder.folder.canSave}>
                  <i className="fa fa-trash" /> 删除
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    );
  }
}
