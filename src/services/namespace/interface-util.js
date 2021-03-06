// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

var jsonStableStringify = require('json-stable-stringify');

module.exports = {
  getMethodData: getMethodData,
  hashInterface: hashInterface
};

/*
 * Given a service interface and method name, retrieve the method signature.
 * @param {interface} interface The service interface
 * @param {string} methodName The name of the method.
 * @return {methodSig} An object containing info about the method.
 */
function getMethodData(interface, methodName) {
  for (var i = 0; i < interface.methods.length; i++) {
    var methodData = interface.methods[i];
    if (methodData.name === methodName) {
      return methodData;
    }
  }
  return null;
}

/*
 * Given a service signature, compute a reasonable hash that uniquely identifies
 * a service without containing unnecessary information.
 * This heuristic comes extremely close to uniquely identifying service methods.
 */
function hashInterface(interface) {
  var pkgPath = interface.pkgPath + '.' + interface.name;
  var methods = interface.methods.map(function(methodData) {
    var inArgList = methodData.inArgs || [];
    return {
      name: methodData.name,
      inArgs: inArgList.map(function(inArg) {
        return inArg.name;
      })
    };
  });
  var output = jsonStableStringify([pkgPath, methods]);
  return output;
}