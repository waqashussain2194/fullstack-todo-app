import React from "react";
import "./styles/Addbar.css";

class AddBar extends React.Component {
  addItem = event => {
    if (event.key === "Enter" && event.target.value.trim() !== "") {
      this.props.addItem(event.target.value);
      event.target.value = ""; // Clear the input field after adding
    }
  };

  render() {
    return (
      <div className="AddBar">
        <input
          className="AddBar-Text"
          type="text"
          placeholder="Enter TODO Item"
          onKeyDown={this.addItem}
        />
      </div>
    );
  }
}

export default AddBar;
