//Player is a singleton 
Player.prototype = new GameObject()
function Player(map, name) {
	if( Player.instance ) 
		return Player.instance
	Player.instance = this
	//set constructor stuff here!
	pos = (Math.floor(Math.random * map.x) + 1, Math.floor(Math.random * map.y) + 1) 
	GameObject.call(this, pos)
	this.name = name
	this.guns = [new Handgun()] //one of each gun, starts with handgun

}

Player.prototype.move=function(input) {
	this.acceleration = input
}

