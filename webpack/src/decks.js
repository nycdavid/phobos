import React from "React";

class Decks extends React.Component {
	constructor(props) {
		super(props);
		console.log(props);
		this.state = {
			decks: props.decks
		};
	}

	render() {
		const rows = this.state.decks.map(deck => {
			return (
				<tr>
					<td>{deck.id}</td>
					<td>{deck.name}</td>
				</tr>
			)
		});

		return (
			<table>
				<thead>
					<th>id</th>
					<th>name</th>
				</thead>
				<tbody>{rows}</tbody>
			</table>
		);
	}
}

export default Decks;
