import React from "react";

class NewCardModal extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      tabGroups: {
        front: 0,
        back: 0,
      },
    };
  }

  activeClass(tabGroup, idx) {
    const index = this.state.tabGroups[tabGroup];
    if (index === idx) {
      return "active";
    }
    return "";
  }

  changeTab(tabGroup, idx) {
    const newTabGroup = {};
    newTabGroup[tabGroup] = idx;
    const tabGroups = { ...this.state.tabGroups, ...newTabGroup };
    this.setState({ tabGroups: tabGroups });
  }

  frontSection() {
    return (
      <div className="card-front">
        <ul className="nav nav-tabs">
          <li className="nav-item">
            <a
              className={`nav-link ${this.activeClass("front", 0)}`}
              href="#"
              onClick={() => this.changeTab("front", 0)}
            >
              Edit
            </a>
          </li>
          <li className="nav-item">
            <a
              className={`nav-link ${this.activeClass("front", 1)}`}
              href="#"
              onClick={() => this.changeTab("front", 1)}
            >
              Preview
            </a>
          </li>
        </ul>
        <form>
          <div className="form-group">
            <label htmlFor="" className="col-form-label">
              Front
            </label>
            <textarea type="text" className="form-control" />
          </div>
        </form>
      </div>
    );
  }

  backSection() {
    return (
      <div className="card-back">
        <ul className="nav nav-tabs">
          <li className="nav-item">
            <a
              className={`nav-link ${this.activeClass("back", 0)}`}
              href="#"
              onClick={() => this.changeTab("back", 0)}
            >
              Edit
            </a>
          </li>
          <li className="nav-item">
            <a
              className={`nav-link ${this.activeClass("back", 1)}`}
              href="#"
              onClick={() => this.changeTab("back", 1)}
            >
              Preview
            </a>
          </li>
        </ul>
        <form>
          <div className="form-group">
            <label htmlFor="" className="col-form-label">
              Back
            </label>
            <textarea type="text" className="form-control" />
          </div>
        </form>
      </div>
    );
  }

  render() {
    return (
      <div
        className="modal fade"
        id="new-card-modal"
        role="dialog"
        aria-hidden="true"
      >
        <div className="modal-dialog" role="document">
          <div className="modal-content">
            <div className="modal-header">
              <h5>New card</h5>
            </div>
            <div className="modal-body">
              {this.frontSection()}
              {this.backSection()}
            </div>
          </div>
        </div>
      </div>
    );
  }
}

export default NewCardModal;
