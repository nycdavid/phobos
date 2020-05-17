import React from "React";
import ReactDOM from "react-dom";
import axios from "axios";

class Decks extends React.Component {
	constructor(props) {
		super(props);
		this.state = { decks: [] };
	}

	render() {
		return null;
	}
}

ReactDOM.hydrate(<Decks />, document.getElementById("app"));
