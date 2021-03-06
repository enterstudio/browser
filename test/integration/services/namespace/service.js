// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

var test = require('prova');
var mercury = require('mercury');
var _ = require('lodash');
var proxyquire = require('proxyquireify')(require);
var mockLRUCache = require('./mocks/lru-cache');
var sampleWorld = require('../../../../src/services/sample-world');

test.timeout(60 * 1000);

// @noCallThru ensures this completely overrdies the original config
// instead of inheriting the properties that are not defined here from
// the original dependency
var vanadiumConfigForTest = {
  '@noCallThru': true
};

// The V23_NAMESPACE environment variable is set by servicerunner in the "make
// test" target.  That environment variable is picked up by the "envify" prova
// transform and used to set process.env.V23_NAMESPACE.
var globalRoot = process.env.V23_NAMESPACE;

// Require namespaceService but using test specific mocks and configs
var namespaceService =
  proxyquire('../../../../src/services/namespace/service.js', {
    '../../vanadium-config': vanadiumConfigForTest,
    'lru-cache': function() {
      return mockLRUCache;
    },
    'vanadium': require('vanadium'),
  });

var sw;

function initSampleWorld(t) {
  if (!sw) {
    var prefix = '';
    sw = sampleWorld.create(prefix);
  }

  return sw.catch(function(err) {
    t.fail(err);
    t.end();
    // Throw again so the promise is not resolved and test does not continue
    // to run.
    throw err;
  });
}

test('getChildren of default namespace root', function(t) {
  initSampleWorld(t).then(run);

  function run() {
    namespaceService.getChildren().
    then(function assertResult(result) {
      assertIsImmutable(t, result);
      // Wait until we finish, we expect 2 top level items: cottage, house
      result.events.on('end', function validate() {
        mercury.watch(result, function(children) {
          assertCottage(children[0]);
          assertHouse(children[1]);
          t.end();
        });
      });
      result.events.on('globError', function(error) {
        t.notOk(error, 'did not expect any globs errors');
        t.end();
      });
    }).catch(t.end);

    function assertCottage(item) {
      assertServer(t, item, {
        name: 'cottage',
        objectName: 'cottage',
        isLeaf: false,
        isMounttable: true
      });
    }

    function assertHouse(item) {
      assertServer(t, item, {
        name: 'house',
        objectName: 'house',
        isLeaf: false,
        isMounttable: true
      });
    }
  }
});

test('getChildren of cottage/lawn', function(t) {
  initSampleWorld(t).then(run);

  function run() {
    namespaceService.getChildren('cottage/lawn').
    then(function assertResult(result) {
      assertIsImmutable(t, result);
      // Wait until we finish, we expect 3 items back, front and
      // master-sprinkler
      result.events.on('end', function validate() {
        mercury.watch(result, function(children) {
          assertBack(children[0]);
          assertSprinkler(children[2]);
          t.end();
        });
      });
      result.events.on('globError', function(error) {
        t.notOk(error, 'did not expect any globs errors');
        t.end();
      });
    }).catch(t.end);

    function assertSprinkler(item) {
      assertServer(t, item, {
        name: 'master-sprinkler',
        objectName: 'cottage/lawn/master-sprinkler',
        isLeaf: true
      });
    }

    function assertBack(item) {
      assertSubtableName(t, item, {
        name: 'back',
        objectName: 'cottage/lawn/back'
      });
    }
  }
});

test('getChildren of rooted ' + globalRoot + '/house/kitchen', function(t) {
  initSampleWorld(t).then(run);

  function run() {
    namespaceService.getChildren(globalRoot + '/house/kitchen').
    then(function assertResult(result) {
      assertIsImmutable(t, result);
      // Wait until we finish, we expect 2 items, lights
      // and smoke-detector
      result.events.on('end', function validate() {
        mercury.watch(result, function(children) {
          assertLightSwitch(children[0]);
          assertSmokeDetector(children[1]);
          t.end();
        });
      });
      result.events.on('globError', function(error) {
        t.notOk(error, 'did not expect any globs errors');
        t.end();
      });
    }).catch(t.end);

    function assertLightSwitch(item) {
      assertServer(t, item, {
        name: 'lights',
        objectName: globalRoot + '/house/kitchen/lights',
        isLeaf: true
      });
    }

    function assertSmokeDetector(item) {
      assertServer(t, item, {
        name: 'smoke-detector',
        objectName: globalRoot + '/house/kitchen/smoke-detector',
        isLeaf: true
      });
    }
  }
});

// The HOUSE_MOUNTTABLE environment variable is set by run-tests.sh.  That
// environment variable is picked up by the "envify" prova transform and used to
// set process.env.HOUSE_MOUNTTABLE.
var hostPortRoot = process.env.HOUSE_MOUNTTABLE;

test('getChildren of rooted ' + hostPortRoot + '/kitchen', function(t) {
  initSampleWorld(t).then(run);

  function run() {
    namespaceService.getChildren(hostPortRoot + '/kitchen').
    then(function assertResult(result) {
      assertIsImmutable(t, result);
      // Wait until we finish, we expect 2 items, lights
      // and smoke-detector
      result.events.on('end', function validate() {
        mercury.watch(result, function(children) {
          assertLightSwitch(children[0]);
          assertSmokeDetector(children[1]);
          t.end();
        });
      });
      result.events.on('globError', function(error) {
        t.notOk(error, 'did not expect any globs errors');
        t.end();
      });
    }).catch(t.end);

    function assertLightSwitch(item) {
      assertServer(t, item, {
        name: 'lights',
        objectName: hostPortRoot + '/kitchen/lights',
        isLeaf: true
      });
    }

    function assertSmokeDetector(item) {
      assertServer(t, item, {
        name: 'smoke-detector',
        objectName: hostPortRoot + '/kitchen/smoke-detector',
        isLeaf: true
      });
    }
  }
});

test('getChildren of non-existing mounttable', function(t) {
  initSampleWorld(t).then(run);

  function run() {
    // TODO(aghassemi) why does namespace library return empty results instead
    // of error when globbing rooted names that don't exist?
    namespaceService.getChildren('/DoesNotExist:666/What/Ever').
    then(function assertResult(result) {
      result.events.on('end', function validate() {
        // Expect empty results
        mercury.watch(result, function(children) {
          t.deepEqual(children, []);
          t.end();
        });
      });
      result.events.on('globError', function(error) {
        // we do actually expect a glob error in this case
        t.ok(error);
      });
    }).catch(t.end);
  }
});

test('getNamespaceItem of leaf server', function(t) {
  initSampleWorld(t).then(run);

  function run() {
    namespaceService.getNamespaceItem('cottage/lawn/master-sprinkler').
    then(function assertItem(itemObs) {
      assertIsImmutable(t, itemObs);
      var item = itemObs();
      assertServer(t, item, {
        name: 'master-sprinkler',
        objectName: 'cottage/lawn/master-sprinkler',
        isLeaf: true
      });
      t.end();
    }).catch(t.end);
  }
});

test('getNamespaceItem of subtable', function(t) {
  initSampleWorld(t).then(run);

  function run() {
    namespaceService.getNamespaceItem('cottage/lawn/back').
    then(function assertItem(itemObs) {
      assertIsImmutable(t, itemObs);
      var item = itemObs();
      assertSubtableName(t, item, {
        name: 'back',
        objectName: 'cottage/lawn/back'
      });
      t.end();
    }).catch(t.end);
  }
});

test('getNamespaceItem of mounttable leaf server', function(t) {
  initSampleWorld(t).then(run);

  function run() {
    namespaceService.getNamespaceItem('cottage').
    then(function assertItem(itemObs) {
      assertIsImmutable(t, itemObs);
      var item = itemObs();
      assertServer(t, item, {
        name: 'cottage',
        objectName: 'cottage',
        isLeaf: false,
        isMounttable: true
      });
      t.end();
    }).catch(t.end);
  }
});

test('search uses caching', function(t) {
  initSampleWorld(t).then(run);

  function run() {
    mockLRUCache.reset();

    namespaceService.search('house', '*').
    then(function assertNoCacheHit() {
      t.notOk(mockLRUCache.wasCacheHit('glob|house/*'),
        'first glob call is not a cache hit');

      // Call second time, there should have been a cache hit
      return namespaceService.search('house', '*');
    }).then(function assertCacheHit() {
      t.ok(mockLRUCache.wasCacheHit('glob|house/*'),
        'second glob call is a cache hit');

      // Call glob with same name, different query
      return namespaceService.search('house', 'foo*');
    }).then(function assertNoCacheHit() {
      t.notOk(mockLRUCache.wasCacheHit('glob|house/foo*'),
        'third glob call with different query is not a cache hit');
      t.end();
    }).catch(t.end);
  }
});

test('getSignature uses caching', function(t) {
  initSampleWorld(t).then(run);

  function run() {
    mockLRUCache.reset();

    namespaceService.getSignature('house/alarm').then(function() {
      t.notOk(mockLRUCache.wasCacheHit('getSignature|house/alarm'),
        'first getSignature call is not a cache hit');
      // Call a second time
      return namespaceService.getSignature('house/alarm');
    }).then(function() {
      t.ok(mockLRUCache.wasCacheHit('getSignature|house/alarm'),
        'second getSignature call is a cache hit');
      // Call a different name
      return namespaceService.getSignature('house/kitchen/smoke-detector');
    }).then(function() {
      t.notOk(mockLRUCache.wasCacheHit(
        'getSignature|house/kitchen/smoke-detector'
      ), 'third getSignature call to a different name is not a cache hit');
      t.end();
    }).catch(t.end);
  }
});

test('test clearCache', function(t) {
  initSampleWorld(t).then(run);

  function run() {
    mockLRUCache.reset();

    // cache bunch of things under house first, then clear cache and validate
    Promise.all([
        namespaceService.getSignature('house/alarm'),
        namespaceService.search('house', '*'),
        namespaceService.search('house', ''),
        namespaceService.search('house-2', ''),
        namespaceService.getRemoteBlessings('house/alarm')
      ])
      .then(function verifyItemsAreCached() {
        t.ok(mockLRUCache.has('getSignature|house/alarm'),
          'signature is cached');
        t.ok(mockLRUCache.has('glob|house/*'),
          'glob is cached');
        t.ok(mockLRUCache.has('glob|house'),
          'glob is cached');
        t.ok(mockLRUCache.has('glob|house-2'),
          'glob is cached');
        t.ok(mockLRUCache.has('getRemoteBlessings|house/alarm'),
          'getRemoteBlessings is cached');
      })
      .then(function clearCacheForHouse() {
        namespaceService.clearCache('house');
      }).then(function verifyCacheWasCleared() {
        t.notOk(mockLRUCache.has('getSignature|house/alarm'),
          'signature cache cleared');
        t.notOk(mockLRUCache.has('glob|house/*'),
          'glob cache cleared');
        t.notOk(mockLRUCache.has('glob|house'),
          'glob cache cleared');
        t.notOk(mockLRUCache.has('getRemoteBlessings|house/alarm'),
          'getRemoteBlessings cache cleared');
        t.ok(mockLRUCache.has('glob|house-2'),
          'house-2 should not be impacted');
        t.end();
      }).then(function deleteEverything() {
        return namespaceService.clearCache();
      }).then(function verifyAllIsGone() {
        t.notOk(mockLRUCache.has('glob|house-2'),
          'house-2 should be gone now');
      }).catch(t.end);
  }
});

// Make RPC: good inputs => no error
var okRPCs = {
  'no input': ['house/alarm', 'status', []],
  'bool input': ['house/living-room/lights', 'flipSwitch', [true]],
  'int input': ['cottage/smoke-detector', 'sensitivity', [2]],
  'float input': ['house/alarm', 'delayArm', [2.5]],
  'string input': ['cottage/pool/speaker', 'playSong', ['Happy Birthday']],
  'slice input': ['house/master-bedroom/speaker', 'addSongs', [
    ['A', 'B']
  ]],
  '2+ inputs': ['cottage/pool/heater', 'start', [70, 5]],
};

_.forOwn(okRPCs, function run(params, inputType) {
  test(
    'makeRPC accepts good input - ' + inputType,
    testMakeRPCNoError.bind(null, params)
  );
});

// Make RPC: bad inputs => error
var badRPCs = {
  //TODO(aghassemi) re-enable after #483
  //'no service':    ['mansion/smoke-detector', 'status', []],
  'no method': ['cottage/pool/speaker', 'status', []],
  'no input': ['cottage/lights', 'flipSwitch', null],
  'bad type': ['cottage/lights', 'flipSwitch', ['notBool']],
  'lacks input': ['cottage/pool/heater', 'start', [80]],
  'invalid input': ['house/living-room/blast-speaker', 'playSong', ['notThere']]
};

_.forOwn(badRPCs, function run(params, inputType) {
  test(
    'makeRPC errors on bad input - ' + inputType,
    testMakeRPCHasError.bind(null, params)
  );
});

// Make RPC: outputs have the expected # of outputs
test('makeRPC returns output properly', function(t) {
  initSampleWorld(t).then(run);

  function run() {
    namespaceService.makeRPC('cottage/alarm', 'panic', []).then(
      function got0Outputs(res) { // 0 outputs: has no result.
        t.ok(res === undefined, '0 outputs => is undefined');

        return namespaceService.makeRPC('house/alarm', 'status', []);
      }
    ).then( // 1 output: (Non-array/slice output) is not an Array.
      function got1Output(res) {
        t.notOk(res instanceof Array, '1 output => not an Array');

        return namespaceService.makeRPC('cottage/smoke-detector', 'test', []);
      }
    ).then( // 1 output: Delayed return. Also not an array.
      function got1OutputDelayed(res) {
        t.notOk(res instanceof Array, '1 output => not an Array');

        return namespaceService.makeRPC('cottage/pool/heater', 'status', []);
      }
    ).then( // 2 outputs: Is an Array of the correct length.
      function got2Outputs(res) {
        var ok = res instanceof Array && res.length === 2;
        t.ok(ok, '2 outputs => length 2 Array');
        t.end();
      }
    ).catch(t.end);
  }
});


/*
 * Runs a test to ensure the makeRPC call terminates without error.
 */
function testMakeRPCNoError(args, t) {
  initSampleWorld(t).then(run);

  function run() {
    namespaceService.makeRPC.apply(null, args).then(function(result) {
      t.pass('completed without error');
      t.end();
    }).catch(function(err) {
      t.end(err);
    });
  }
}

/*
 * Runs a test to ensure the makeRPC call terminates with an error.
 */
function testMakeRPCHasError(args, t) {
  initSampleWorld(t).then(run);

  function run() {
    namespaceService.makeRPC.apply(null, args).then(function(result) {
      t.fail('should not have completed without error');
      t.end();
    }).catch(function(err) {
      t.pass('correctly returned an error');
      t.end();
    });
  }
}

/*
 * Test helpers
 */
function assertServer(t, item, vals) {
  assertMountedName(t, item, vals.name);
  assertObjectName(t, item, vals.objectName);
  assertIsServer(t, item);
  if (vals.isLeaf === true) {
    assertIsLeaf(t, item);
  } else if (vals.isLeaf === false) {
    assertIsNotLeaf(t, item);
  }

  t.equals(!!item.isMounttable, !!vals.isMounttable,
    item.mountedName + ': is a mounttable');
}

function assertSubtableName(t, item, vals) {
  assertMountedName(t, item, vals.name);
  assertObjectName(t, item, vals.objectName);
  assertIsNotLeaf(t, item);
  t.ok(item.hasMountPoint, ': is subtable node');
}

function assertMountedName(t, item, val) {
  t.ok(item.mountedName, item.mountedName + ': has a mounted name');
  t.equal(item.mountedName, val, item.mountedName + ': mounted name matches');
}

function assertObjectName(t, item, val) {
  t.ok(item.objectName, item.mountedName + ': has an object name');
  t.equal(item.objectName, val, item.mountedName + ': object name matches');
}

function assertIsServer(t, item) {
  t.ok(item.hasServer, item.mountedName + ': is a server');
}

function assertIsLeaf(t, item) {
  t.equal(item.isLeaf, true, item.mountedName + ': is leaf');
}

function assertIsNotLeaf(t, item) {
  t.equal(item.isLeaf, false, item.mountedName + ': is not leaf');
}

function assertIsImmutable(t, observable) {
  t.ok(observable.set === undefined, 'is immutable');
}

//TODO(aghassemi)
//Tests for:
// Recursive glob
// Glob with some keyword
// Ensuring array is updated when nodes get mounted and unmounted (when we use
// watchGlob)
// getEndpoints()
// getPermissions()
// resolveToMounttable()
// getRemoteBlessings()