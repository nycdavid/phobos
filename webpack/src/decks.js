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
		const rows = this.state.decks.map((deck, i) => {
			return (
				<tr key={i}>
					<td>{deck.id}</td>
					<td>{deck.name}</td>
					<td>
						<a href="#">New Card</a>
						<a href="#">Study</a>
					</td>
				</tr>
			)
		});

		return (
			<table className="table">
				<thead>
					<tr>
						<th>ID</th>
						<th>Name</th>
						<th>Actions</th>
					</tr>
				</thead>
				<tbody>{rows}</tbody>
			</table>
		);
	}
}

export default Decks;
