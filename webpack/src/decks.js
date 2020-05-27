import React from "React";

import NewCardModal from "./new_card_modal";


class Decks extends React.Component {
  constructor(props) {
    super(props);
    console.log(props);
    this.state = {
      decks: props.decks,
			modalOpen: false
    };
  }

	openModal() {
		this.setState({ modalOpen: true });
	}

  render() {
    const rows = this.state.decks.map((deck, i) => {
      return (
        <tr key={i}>
          <td>{deck.id}</td>
          <td>{deck.name}</td>
          <td>
            <a
              data-target="#new-card-modal"
              data-toggle="modal"
              href="#"
							onClick={this.openModal.bind(this)}
            >
              New Card
            </a>
            <a href="#">Study</a>
          </td>
        </tr>
      );
    });

    return (
      <div>
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
        <NewCardModal modalOpen={this.state.modalOpen} />
      </div>
    );
  }
}

export default Decks;
