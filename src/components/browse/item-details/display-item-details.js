var browseService = require('../../../services/browse-service');
var smartService = require('../../../services/smart-service');
var debug = require('debug')('display-item-details');

module.exports = displayItemDetails;

/*
 * Ask the browseService for a service signature.
 * Use the signature and smartService to pick which RPCs to do automatically.
 */
function displayItemDetails(state, data) {
  var name = data.name;

  // Don't refresh if we are already looking at this name's details.
  if (state.itemName() === name) {
    return;
  }
  // Log the URL to the smart service as a potential shortcut.
  smartService.record('learner-shortcut', {name: name});

  // Save every time we look at a service's details.
  smartService.save('learner-shortcut');

  // Set the new name and reset the selected method and outputs.
  // TODO(alexfandrianto): Instead of resetting, should we remember this info?
  state.itemName.set(name);
  state.selectedMethod.set('');

  // TODO(aghassemi)
  // any better way than splice to tell Mercury all of array changed?
  state.methodOutputs.splice(0, state.methodOutputs.getLength());

  browseService.signature(name).then(function(signatureResult) {
    state.signature.set(signatureResult);

    // Go through each signature method and decide whether to perform it or not.
    for (var m in signatureResult) {
      if (signatureResult.hasOwnProperty(m)) {
        var param = signatureResult[m];
        var input = {
          name: name,
          methodName: m,
          signature: signatureResult,
          hasParams: param.inArgs.length !== 0,
        };

        var details = state.details.get(data.name);

        // We will neither recommend methods that take input parameters nor
        // re-recommend a method.
        if (input.hasParams || (details && details[m] !== undefined)) {
          continue;
        }

        // If the prediction power is strong enough, recommend the method.
        var prediction = smartService.predict('learner-autorpc', input);
        if (prediction > 0.5) {
          debug('Recommend', m, 'with', prediction);

          // Set the state detail with the prediction value (a float).
          var detail = state.details.get(data.name);
          if (detail === undefined) {
            detail = {};
          }
          detail[input.methodName] = prediction;
          state.details.put(input.name, detail);
        }
      }
    }
  }, function(err) {
    debug('Failed to get signature',
      name,
      err, (err && err.stack) ? err.stack : undefined
    );
    state.signature.set('');
  });
}