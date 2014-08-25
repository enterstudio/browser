var mercury = require('mercury');
var AttributeHook = require('../../../lib/mercury/attribute-hook');
var insertCss = require('insert-css');
var displayItemDetails = require('./display-item-details');
var makeRPC = require('./make-rpc');
var browseService = require('../../../services/browse-service');
var h = mercury.h;
var css = require('./index.css');

module.exports = create;
module.exports.render = render;

/*
 * ItemDetails component provides user interfaces for displaying details for
 * a browse item such is its type, signature, etc.
 */
function create() {

  var state = mercury.struct({
    /*
     * Item name to display settings for
     * @type {string}
     */
    itemName: mercury.value(''),

    /*
     * Method signature for the name, if pointing to a server
     * @type {Object}
     */
    signature: mercury.value(null),

    selectedTabIndex: mercury.value(0),

    details: mercury.varhash(),
  });

  var events = mercury.input([
    'displayItemDetails',
    'tabSelected',
    'methodSelected',
  ]);

  wireUpEvents(state, events);

  return {
    state: state,
    events: events
  };
}

function render(state, events) {
  insertCss(css);
  return [h('paper-tabs.tabs', {
      'selected': new AttributeHook(state.selectedTabIndex),
      'noink': new AttributeHook(true)
    }, [
      h('paper-tab.tab', {
        'ev-click': mercury.event(events.tabSelected, {
          index: 0
        })
      }, 'Details'),
      h('paper-tab.tab', {
        'ev-click': mercury.event(events.tabSelected, {
          index: 1
        })
      }, 'Signature'),
    ]),
    h('core-selector', {
      'selected': new AttributeHook(state.selectedTabIndex)
    }, [
      h('div.tab-content', renderDetailsTab()),
      h('div.tab-content', renderSignatureTab())
    ])
  ];

  function renderDetailsTab() {
    var typeInfo = browseService.getTypeInfo(state.signature);
    var displayItems = [
      renderFieldItem('Name', (state.itemName || '<root>')),
      renderFieldItem('Type', typeInfo.name, typeInfo.description),
    ];

    // In addition to the Name and Type, render additional service details.
    var details = state.details;
    for (var method in details[state.itemName]) {
      if (details[state.itemName].hasOwnProperty(method)) {
        displayItems.push(
          renderFieldItem(
            method,
            details[state.itemName][method]
          )
        );
      }
    }

    return [
      h('div', displayItems),
    ];
  }

  function renderSignatureTab() {
    var methods = [];
    var sig = state.signature;
    for (var m in sig) {
      if (sig.hasOwnProperty(m)) {
        methods.push(renderMethod(m, sig[m]));
      }
    }

    if (methods.length > 0) {
      return h('div.signature', methods);
    } else {
      return h('div.empty', 'No signature');
    }

    function renderMethod(name, param) {
      var text = name + '(';
      for (var i = 0; i < param.inArgs.length; i++) {
        var arg = param.inArgs[i];
        if (i > 0) {
          text += ',';
        }
        text += arg;
      }
      text += ')';
      if (param.isStreaming) {
        text += ' - streaming';
      }
      return h('pre', {
        'ev-click': mercury.event(events.methodSelected, {
          name: state.itemName,
          methodName: name,
          signature: sig,
          hasParams: param.inArgs.length !== 0,
        })
      }, text);
    }
  }
}

/*TODO(aghassemi) make a web component for this*/
function renderFieldItem(label, content, tooltip) {

  var hlabel = h('h4', label);
  if (tooltip) {
    // If there is a tooltip, wrap the content in it
    content = h('core-tooltip.tooltip', {
      'label': new AttributeHook(tooltip),
      'position': 'right',
    }, content);
  }

  return h('div.field', [
    h('h4', hlabel),
    h('div.content', content)
  ]);
}

// Wire up events that we know how to handle
function wireUpEvents(state, events) {
  events.displayItemDetails(displayItemDetails.bind(null, state));
  events.tabSelected(function(data) {
    state.selectedTabIndex.set(data.index);
  });
  events.methodSelected(makeRPC.bind(null, state));
}