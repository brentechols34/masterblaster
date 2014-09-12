/** @jsx React.DOM */ 
define([
'jquery',
'react'
],

function ($, React) {


	var GameView = React.createClass({
		getInitialState: function () {
			return {};
		},

		render: function () {
			return (<div id="gameWindow">
					"This is a game window!"
				</div>);
		}
	});


	return GameView;
})