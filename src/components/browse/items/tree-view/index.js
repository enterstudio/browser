var mercury = require('mercury');
var extend = require('extend');

var polymerEvent = require('../../../../lib/mercury/polymer-event');
var loadChildren = require('./load-children');
var getServiceIcon = require('../../get-service-icon');

var h = mercury.h;

module.exports = create;
module.exports.render = render;
module.exports.loadChildren = loadChildren;

function create() {

  var state = mercury.varhash({
    /*
     * Map of objectNames to child namespace items
     * @see services/namespace/item
     * @type {varhash<string,Array<namespaceitem>>}
     * Always contains the children for every item in the tree (including
     *   leaf nodes) so we can know whether to display the expand icon
     */
    childrenMap: mercury.varhash({}),

    /*
     * The current item to be used as the root of the tree
     * @see services/namespace/item
     * @type {namespaceitem}
     */
    rootItem: mercury.value(null)
  });

  var events = mercury.input([
    'openChange', // expand / collapse of tree node
    'activate'    // tap on tree node
  ]);

  wireUpEvents(state, events);

  return {
    state: state,
    events: events
  };
}

function render(state, events, browseState, browseEvents) {
  var item = state.rootItem;  // start at the root
  if (item === null) { return; }  // TODO(wm) Maybe show "Loading..."?

  var rootEvents = { // events to attach to root of tree
    'ev-openchange': polymerEvent(events.openChange),
    'ev-activate': polymerEvent(
      events.activate,
      { browseEvents: browseEvents }
    )
  };
  return h('div#tree-container', [ h('h2', 'Tree View'),
      createTreeNode(state, browseState.selectedItemName, item, rootEvents)
      ]);
}

/*
 * Recursively render tree from the bottom up
 * Has to be from bottom up because virtual DOM is immutable.
 * events is extra properties to add (used for root events)
 */
function createTreeNode(state, selected, item, extraprops) {
  var childrenArr = state.childrenMap[item.objectName];
  var descendants = []; // all viewed descendants of this item
  if (childrenArr) {
    descendants = childrenArr.map(function(child) {
        return createTreeNode(state, selected, child);
    });
  }
  var props = {
    attributes: {
      label: item.mountedName || '<root>',
      icon:  getServiceIcon(item.isServer ? item.serverInfo.typeInfo.key : ''),
      itemTitle: item.objectName,
      highlight: (item.objectName === selected)
    },
    objectName: item.objectName
  };
  if (extraprops) { // root of tree
    extend(props, extraprops);
  }
  return h('tree-node', props, descendants);
}

function wireUpEvents(state, events) {
  events.openChange(function(data) {
    var objectName = data.polymerDetail.node.objectName;
    // load up the immediate children of a newly displayed item
    // so we know if that item can be expanded
    state.childrenMap[objectName].map(function(child) {
      loadChildren(state, { parentName: child.objectName() });
    });
  });

  events.activate(function(data) {
    var objectName = data.polymerDetail.node.objectName;
    data.browseEvents.selectItem({
      name: objectName
    });
  });
}