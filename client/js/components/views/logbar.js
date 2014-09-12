/** @jsx React.DOM */
define([
'jquery',
'react',
'js/components/controls/keyCapture.js'
],

function ($, React, keylogger) {

	var KeyList = React.createClass({

		updateData: function(key) {
			this.setState({data: key})
		},

		getInitialState: function() {
			return {data: null}
		},

		componentDidMount: function() {
			keylogger.addListener(this.updateData)
		},

		componentWillUnmount: function() {
			keylogger.removeListener(this.updateData)
		},

		render: function() {
			return <div id="keyList">
						{this.state.data}
				   </div>;
		}
	});

	return KeyList;
})