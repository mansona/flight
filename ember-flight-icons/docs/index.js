import Component from '@glimmer/component';
import { tracked } from '@glimmer/tracking';
import { action } from '@ember/object';
import { getOwner } from '@ember/application';

const defaultSize = '24';

const checkIsShown = function (searchText, meta) {
  if (searchText === '') {
    return true;
  }

  if (meta.searchable.indexOf(searchText) !== -1) {
    return true;
  }

  return false;
};

export default class IndexComponent extends Component {
  get contextRootURL() {
    const config = getOwner(this).resolveRegistration('config:environment');
    return config.rootURL || '/';
  }

  @tracked model;

  constructor() {
    super(...arguments);
    fetch(
      `${this.contextRootURL}@hashicorp/ember-flight-icons/icons/_catalog.json`
    ).then(async (response) => {
      const json = await response.json();

      this.model = json.map(({ Name, Size }) => {
        return {
          name: `${Name}`,
          size: `${Size}`,
          searchable: `${Name}`,
        };
      });
    });
  }

  @tracked selectedIcon = 'auto-apply';
  @tracked size = '24';
  @tracked color = 'currentColor';
  @tracked searchText = '';

  @tracked search;
  @tracked emptyResults = false;

  get iconHbsCode() {
    let iconHbsCode = `<FlightIcon @name="${this.selectedIcon}"`;

    if (this.size !== defaultSize) {
      iconHbsCode += ` @size=${this.size}`;
    }

    if (this.color) {
      iconHbsCode += ` @color="${this.color}"`;
    }

    iconHbsCode += '/>';

    return iconHbsCode;
  }

  @action
  async updateSearchText(value, signal) {
    await new Promise((resolve) => setTimeout(resolve, 190));

    if (signal.aborted) {
      return;
    }

    this.search = value;
    const lowcased = value.toLowerCase();

    for (let i = 0; i < this.model.length; i++) {
      const item = this.model[i];

      set(item, 'isHidden', !checkIsShown(lowcased, item));
    }

    this.emptyResults = this.model.every(({ isHidden }) => isHidden);
  }

  @action
  debouncedUpdate(event) {
    if (this.ctrl) {
      this.ctrl.abort();
    }
    this.ctrl = new AbortController();
    this.updateSearchText(event.target.value, this.ctrl.signal);
  }

  @action
  updateSelectedItem(event) {
    const iconWrapper = event.target.closest('.demo-icon');
    if (iconWrapper && iconWrapper.dataset.Name) {
      this.selectedIcon = iconWrapper.dataset.Name;
    }
  }
}
