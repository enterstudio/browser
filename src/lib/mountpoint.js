//TODO(aghassemi) move to core Veyron API?

module.exports = MountPoint;

/**
 * MountPoint handles manipulating and querying a namespace
 * @param {object} rt A veyron runtime.
 * @param {object} namespace A veyron namespace instance.
 * @param {...string} addressParts Parts of the address to join
 * @constructor
 */
function MountPoint(rt, namespace, addressParts) {
  this.rt = rt;
  this.namespace = namespace;
  this.name = Array.prototype.slice.call(arguments, 2).join('/');
  this._terminalNames = null;
}

/**
 * A helper method that returns the terminal names for this
 * MountPoint and memoizes them.
 * @return {Promise} A promise that resolves to a list of terminal names.
 */
MountPoint.prototype._getTerminalNames = function() {
  // We resolve to a terminal name manually because veyron rpc calls
  // wont usually resolve a name if it's to a namespace.  We
  // would like to interact with all kinds of servers.
  if (!this._terminalNames) {
    this._terminalNames = this.namespace.resolveMaximally(this.name);
  }
  return this._terminalNames;
};

/**
 * appendToPath appends to the mountpoint path
 * @param {...string} toAdd strings to add to the path.
 * @return {MountPoint} a new mount point with the path args appended
 * to the current path.
 */
MountPoint.prototype.appendToPath = function(toAdd) {
  var args = Array.prototype.slice.call(arguments);
  if (this.name.length > 0) {
    args.unshift(this.name);
  }
  return new MountPoint(this.rt, this.namespace, args.join('/'));
};

/**
 * mount mounts a target to the current mount point.
 * @param {string} target The target to be mounted.
 * @return {promise} a promise that completes when it is mounted
 */
MountPoint.prototype.mount = function(target) {
  var rt = this.rt;
  return this._getTerminalNames().then(function(terminalNames) {
    // TODO(mattr): We should try all the names instead of just the first.
    // Perhpas the library should allow me to pass a list of names.
    return rt.bindTo(terminalNames[0]).then(function(mtService) {
      return mtService.mount(target, 0);
    });
  });
};

/**
 * glob makes a glob request to a server relative to the current mountpoint.
 * @param {string} expr The glob expression e.g. A/B/*.
 * @return {promise} a promise to a list of results
 */
MountPoint.prototype.glob = function(expr) {
  var results = [];
  var rt = this.rt;
  return this._getTerminalNames().then(function(terminalNames) {
    // TODO(mattr): We should try all the names instead of just the first.
    // Perhpas the library should allow me to pass a list of names.
    return rt.bindTo(terminalNames[0]).then(function(mtService) {
      var promise = mtService.glob(expr);
      var stream = promise.stream;

      stream.on('data', function(val) {
        if (val) {
          results.push(val);
        }
      });

      return promise.then(function() {
        return results;
      });
    });
  });
};
