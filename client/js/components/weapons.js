function Weapon() {
	this.reload = -1
	this.name = "base"
	this.projectile = null
	this.clipsize = 0
	this.currentammo = 0
	this.maxammo = 0
}

Weapon.prototype.fire=function(direction, location) {
	alert("firing weapon " + this.name)
	//probably need to attempt a wait
	new this.projectile(direction, location)
}

Weapon.prototype.reload=function() {
	this.currentammo = this.maxammo
	this.clipsize = this.clipsize - 1
}

Handgun.prototype = new Weapon()
function Handgun() {
	this.name = "handgun"
	this.reload = 2
	this.projectile = Bullet.class
}

MissleLauncher.prototype = new Weapon()
function MissleLauncher() {
	this.name = "misslelauncher"
	this.reload = 10
	this.projectile = Missle.class
}

RailGun.prototype = new Weapon()
function RailGun() {
	this.name = "railgun"
	this.reload = 8
	this.projectile = Rail.class
}
