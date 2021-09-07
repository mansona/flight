# Welcome

Welcome to Field Guide 🎉

You have now successfully setup your own Field Guide, if things are looking a bit bare bones right now that's because Field Guide is designed to provide as few styles as possible so it doesn't interfere with your own styleguide or design system.

If you want to see an example of a full Field Guide in the wild you can [check out the Ember Styleguide](https://ember-styleguide.netlify.com/).


<section class="ds-section">
  <h2 class="ds-h2">
    Available Icon List
  </h2>
  <div class="ds-form-group-inline">
    <label class="ds-input-label" for="ds-input-search">
      Search to live-filter results:
    </label>
    <input
      value={{this.search}}
      placeholder="e.g. arrow"
      autofocus={{false}}
      class="ds-input"
      id="ds-input-search search-field"
      {{on 'input' this.debouncedUpdate}}
    />
  </div>
  <ul class="ds-ul-grid">
    {{#each this.model as |meta|}}
      <li class="ds-li {{if meta.isHidden 'd-none'}}">
        <div class="ds-icon-frame">
          <FlightIcon
            @name={{Onlyname meta.name}}
            @size={{if (eq meta.size '24') '24'}}
            class="demo-icon"
          />
        </div>
        <p>{{meta.name}}</p>
      </li>
    {{/each}}
  </ul>
</section>
