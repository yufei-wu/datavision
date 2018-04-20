import kbn from 'app/core/utils/kbn';
import { Variable, containsVariable, assignModelProperties, variableTypes } from './variable';

export class DatasourceVariable implements Variable {
  regex: any;
  query: string;
  options: any;
  current: any;
  refresh: any;

  defaults = {
    type: 'datasource',
    name: '',
    hide: 0,
    label: '',
    current: {},
    regex: '',
    options: [],
    query: '',
    refresh: 1,
  };

  /** @ngInject **/
  constructor(private model, private datasourceSrv, private variableSrv, private templateSrv) {
    assignModelProperties(this, model, this.defaults);
    this.refresh = 1;
  }

  getSaveModel() {
    assignModelProperties(this.model, this, this.defaults);

    // dont persist options
    this.model.options = [];
    return this.model;
  }

  setValue(option) {
    return this.variableSrv.setOptionAsCurrent(this, option);
  }

  updateOptions() {
    var options = [];
    var sources = this.datasourceSrv.getMetricSources({ skipVariables: true });
    var regex;

    if (this.regex) {
      regex = this.templateSrv.replace(this.regex, null, 'regex');
      regex = kbn.stringToJsRegex(regex);
    }

    for (var i = 0; i < sources.length; i++) {
      var source = sources[i];
      // must match on type
      if (source.meta.id !== this.query) {
        continue;
      }

      if (regex && !regex.exec(source.name)) {
        continue;
      }

      options.push({ text: source.name, value: source.name });
    }

    if (options.length === 0) {
      options.push({ text: '没有任何数据源', value: '' });
    }

    this.options = options;
    return this.variableSrv.validateVariableSelectionState(this);
  }

  dependsOn(variable) {
    if (this.regex) {
      return containsVariable(this.regex, variable.name);
    }
    return false;
  }

  setValueFromUrl(urlValue) {
    return this.variableSrv.setOptionFromUrl(this, urlValue);
  }

  getValueForUrl() {
    return this.current.value;
  }
}

variableTypes['datasource'] = {
  name: 'Datasource',
  ctor: DatasourceVariable,
  description: '使您能够动态切换多个面板的数据源',
};
