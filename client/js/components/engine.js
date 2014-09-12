/* Defines the basic movement of the objects in the game */
definte([],

function() {

	var Engine = {

		updatePosition: function(object) {
			object.x = object.x + object.velocity.x;
			object.y = object.y + object.velocity.y;
		},

		updateVelocity: function(object, accel) {
			object.velocity.x = object.velocity.x + accel.x;
			object.velocity.y = object.velocity.y + accel.y
		},

		overrideVelocity: function(object, newVel) {
			object.velocity = newVel
		},

		overrideObject: function(object, newObj) {
			object = newObj
		}
	};

	return Engine;
})