/** @jsx React.DOM */
define([
'jquery',
'react',
'js/components/views/gameWindow.js',
'js/components/views/sidebar.js',
'js/components/views/logbar.js'
],

function ($, React, Gamewindow, Sidebar, Logbar) {

	var Container = React.createClass({

		//Removed sidebar
		render: function () {
			return (<div id="container">
				<div id="gameContainer">
					<Gamewindow />
				</div>
				<div id="log">
					<Logbar />
				</div>
			</div>);
		}
	})

	return Container;
})