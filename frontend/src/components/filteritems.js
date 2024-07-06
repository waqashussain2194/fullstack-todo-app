import React from "react";
import "./styles/FilterItems.css";

class FilterItems extends React.Component {
  render() {
    const { setFilter, filter } = this.props;

    return (
      <div className="filters">
        <button 
          className={filter === "all" ? "active" : ""} 
          onClick={() => setFilter("all")}
        >
          All
        </button>
        <button 
          className={filter === "notdone" ? "active" : ""} 
          onClick={() => setFilter("notdone")}
        >
          Not Done
        </button>
        <button 
          className={filter === "done" ? "active" : ""} 
          onClick={() => setFilter("done")}
        >
          Done
        </button>
      </div>
    );
  }
}

export default FilterItems;
