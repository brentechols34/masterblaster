/* Captures a key press, generating an action for the player to be sent off */
define([
'jquery',
'react'
],

function ($, React) {

	var key = null;
	var listeners = [];

	var keyLogger = {
		addListener: function(callback) {
			listeners.push(callback);
		},

		notifyListeners: function() {
			listeners.forEach(function (a) {
				a(key);
			});
		},
		removeListener: function(a) {
			listeners = listeners.filter(function(x) {a != x});
		}

	}

	$(document).keypress(function(e) {
		key = e.keyCode;
		keyLogger.notifyListeners();
		//make a moving window of key presses, not all of them
	});

	return keyLogger;	
})