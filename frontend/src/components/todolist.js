import React from "react";
import "./styles/Todolist.css";

class Todolist extends React.Component {
  
  isDone(done) {
    return done ? "Done" : "Not Done";
  }

  createItem(item) {
    return (
      <div className="ListItem" key={item.id} id={item.id}>
        <div className="Title">
          <div className="RemoveItem" onClick={() => this.props.removeItem(item.id)}>
            X
          </div>
          {item.item}
        </div>
        <div className="Status" onClick={() => this.props.toggleDone(item.id)}>
          {this.isDone(item.done)}
        </div>
      </div>
    );
  }

  render() {
    var items = this.props.items;
    return (
      <div className="TodoList">
        <div className="List">{items.map(item => this.createItem(item))}</div>
      </div>
    );
  }
}

export default Todolist;
