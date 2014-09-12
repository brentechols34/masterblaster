/** @jsx React.DOM */ 
define([
'jquery',
'react'
],

function ($, React) {

	console.log("STUFF")

	var Sidebar = React.createClass({
		getInitialState: function () {
			return {};
		},

		render: function () {
			return (<div id="sideBar">
					"sidebar full of fun stuff"
				</div>);
		}
	});

	return Sidebar;
})