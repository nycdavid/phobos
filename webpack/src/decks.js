import React from "React";

import NewCardModal from "./new_card_modal";

class Decks extends React.Component {
  constructor(props) {
    super(props);
    console.log(props);
    this.state = {
      decks: props.decks,
    };
  }

  render() {
    const rows = this.state.decks.map((deck, i) => {
      return (
        <tr key={i}>
          <td>{deck.id}</td>
          <td>{deck.name}</td>
          <td>
            <a data-target="#new-card-modal" data-toggle="modal" href="#">
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
        <NewCardModal />
      </div>
    );
  }
}

export default Decks;
