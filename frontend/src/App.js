import React from "react";

import Header from "./components/header";
import AddBar from "./components/addbar";
import TodoList from "./components/todolist";
import FilterItems from "./components/filteritems";
import "./App.css";

class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      items: [],
      filter: "all",
    };
  }

  removeItem = (id) => {
    fetch(`http://localhost:8081/item/delete/${id}`, {
      method: 'DELETE',
    })
      .then(response => {
        if (response.ok) {
          this.setState({
            items: this.state.items.filter(item => item.id !== id),
          });
        } else {
          console.error('Error deleting item');
        }
      })
      .catch(error => console.error('Error deleting item:', error));
  }

  addItem = (newItemText) => {
    fetch(`http://localhost:8081/item/create`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ item: newItemText })
    })
      .then(response => response.json())
      .then(response => {
        if (response.items && response.items.id) {
          this.setState({ items: [...this.state.items, response.items] });
        } else {
          console.error('Error: New item has no id');
        }
      })
      .catch(error => console.error('Error adding item:', error));
  };

  toggleDone = (id) => {
    let items = [...this.state.items];
    let item = items.find(item => item.id === id);
    item.done = !item.done;

    fetch(`http://localhost:8081/item/update/${id}/${item.done}`, {
      method: 'PATCH',
    })
      .then(response => {
        if (response.ok) {
          this.setState({ items });
        } else {
          console.error('Error updating item');
        }
      })
      .catch(error => console.error('Error updating item:', error));
  }

  fetchItems = (filter) => {
    let url = "http://localhost:8081/items";
    if (filter !== "all") {
      const done = filter === "done" ? "true" : "false";
      url = `http://localhost:8081/items/filter/${done}`;
    }

    fetch(url)
      .then(res => res.json())
      .then(json => {
        if (json.items) {
          this.setState({ items: json.items });
        } else {
          console.error('Error: Items have no ids');
        }
      })
      .catch(error => console.error('Error fetching items:', error));
  }

  componentDidMount() {
    this.fetchItems(this.state.filter);
  }

  setFilter = (filter) => {
    this.setState({ filter }, () => this.fetchItems(filter));
  }

  render() {
    const { items, filter } = this.state;
    return (
      <div className="App">
        <Header />
        <AddBar addItem={this.addItem} />
        <FilterItems 
          setFilter={this.setFilter} 
          filter={filter} 
        />
        <TodoList 
          items={items} 
          removeItem={this.removeItem} 
          toggleDone={this.toggleDone} 
        />
      </div>
    );
  }
}

export default App;
