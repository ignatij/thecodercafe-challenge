# thecodercafe-challenge

# 📦 Max Flow (Edmonds–Karp in Go)

## 📖 Description
This project implements the **Maximum Flow problem** using the **Edmonds–Karp algorithm**, a BFS-based version of the **Ford–Fulkerson method**.  

The maximum flow problem asks:  
> Given a directed graph where each edge has a capacity, and two special nodes called the **source** and **sink**, what is the maximum amount of “flow” that can be sent from the source to the sink without exceeding any edge’s capacity?

A typical example is transporting goods, passengers, or data through a network of roads, pipes, or communication links.

---

## ⚙️ Algorithm
- **Ford–Fulkerson method**: repeatedly find paths from source to sink where flow can still be sent (augmenting paths), push as much as possible along those paths, and update the residual capacities.  
- **Edmonds–Karp algorithm**: a specific implementation of Ford–Fulkerson that always finds augmenting paths with **Breadth-First Search (BFS)**.  
- **Complexity**: `O(VE^2)`, where `V` = vertices, `E` = edges.  

### Key points:
- Uses a **residual graph**: forward edges store remaining capacity, reverse edges store the ability to “undo” or reroute flow.  
- BFS guarantees shortest augmenting paths (fewest edges), avoiding infinite loops and ensuring termination in polynomial time.  

---

## 🛠 Implementation Details
- Written in **Go**.  
- Graph stored as `map[string]map[string]int` (adjacency map of capacities).  
- BFS (`bfsFindPath`) finds augmenting paths and reconstructs them with a parent map.  
- Reverse capacities are updated to allow flow rerouting if a better path is found.  

---

## ▶️ Usage
1. Put your input data in `input.txt` (custom format based on transmissions, alerts, and critical nodes).  
2. Run:
   ```bash
   go run main.go
