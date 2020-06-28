import React from "react";
import marked from "marked";
import axios from "axios";
import Modal from "react-modal";

marked.setOptions({ gfm: true });

class NewCardModal extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      closeModalFn: props.closeModalFn,
      tabGroups: {
        front: 0,
        back: 0,
      },
      content: {
        front: "",
        back: "",
      },
    };
  }

  componentDidUpdate(prevProps, prevState, snapshot) {
    const modalOpen = this.props.modalOpen;

    if (modalOpen !== prevProps.modalOpen) {
      this.setState({ modalOpen: modalOpen });
    }
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
    this.setState({ ...this.state, tabGroups: tabGroups });
  }

  handleChange(evt) {
    this.setState({ content: { front: evt.target.value } });
  }

  handleChangeBack(evt) {
    console.log(this.state);
    const markedimp = marked;
    this.setState({ content: { back: evt.target.value } });
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
        {this.state.tabGroups.front === 0 ? (
          <form>
            <div className="form-group">
              <label htmlFor="" className="col-form-label">
                Front
              </label>
              <textarea
                type="text"
                className="form-control"
                value={this.state.content.front}
                onChange={this.handleChange.bind(this)}
              />
            </div>
          </form>
        ) : (
          <div
            dangerouslySetInnerHTML={{
              __html: marked(this.state.content.front),
            }}
          />
        )}
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

  save() {
    axios
  }

  render() {
    return (
      <Modal isOpen={this.state.modalOpen} ariaHideApp={false}>
        <div className="modal-dialog" role="document">
          <div className="modal-content">
            <div className="modal-header">
              <h5>New card</h5>
                <button className="close" type="button" aria-label="close" onClick={this.props.closeModalFn}>
                  <span aria-hidden="true">
                    &times;
                  </span>
              </button>
            </div>
            <div className="modal-body">
              {this.frontSection()}
              {this.backSection()}
            </div>
            <div className="modal-footer">
              <button
                type="button"
                className="btn btn-primary"
                onClick={this.save.bind(this)}
              >
                Save
              </button>
            </div>
          </div>
        </div>
      </Modal>
    );
  }
}

export default NewCardModal;
