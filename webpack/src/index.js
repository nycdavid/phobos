import React from "React";
import ReactDOM from "react-dom";

class Application extends React.Component {
	render() {
		return (<h1>Application</h1>);
	}
}

ReactDOM.hydrate(<Application />, document.getElementById("app"));
