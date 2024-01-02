import React from "react";

import Header from "./components/header";
import AddBar from "./components/addbar";
import TodoList from "./components/todolist";

import "./App.css";

class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      items: [],
    };
  }

  removeItem = (id) => {
    fetch(`http://localhost:8081/item/delete/${id}`).then(
      this.setState({
        items: this.state.items.filter(item => item.id !== id),
      })
    );
  }

  addItem = (newItemText) => {
    fetch(`http://localhost:8081/item/create/${newItemText}`)
      .then(response => response.json())
      .then(newItem => {
        console.log(newItem);
        this.setState({ items: [...this.state.items, newItem.items] });
      })
      .catch(error => console.error('Error adding item:', error));
  };

  toggleDone = (id) => {
    let items = [...this.state.items];
    let item = items.find(item => item.id === id);
    item.done = !item.done;

    fetch(`http://localhost:8081/item/update/${id}/${item.done}`).then(
      this.setState({ items })
    );
  }

  componentDidMount() {
    fetch("http://localhost:8081/items")
      .then(res => res.json())
      .then(json => this.setState({ items: json.items }));
  }

  render() {
    return (
      <div className="App">
        <Header />
        <AddBar addItem={this.addItem} />
        <TodoList 
          items={this.state.items} 
          removeItem={this.removeItem} 
          toggleDone={this.toggleDone} 
        />
      </div>
    );
  }
}

export default App;
