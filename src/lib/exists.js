// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Given a collection of objects, returns true if all of them exist
 * Returns false as soon as one does not exist.
 * @param {*} [...] objects Objects to check existence of
 * @return {bool} Whether all of the given objects exist or not
 */

module.exports = exists;

function exists(objects) {
  if (!Array.isArray(objects)) {
    objects = [objects];
  }
  for (var i = 0; i < objects.length; i++) {
    var obj = objects[i];
    if (typeof obj === 'undefined' || obj === null) {
      return false;
    }
  }

  return true;
}