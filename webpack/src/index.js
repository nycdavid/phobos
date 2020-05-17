import React from "React";
import ReactDOM from "react-dom";

class Decks extends React.Component {
	render() {
		return (<h1>Deck</h1>)
	}
}

ReactDOM.hydrate(<Decks />, document.getElementById("app"));
