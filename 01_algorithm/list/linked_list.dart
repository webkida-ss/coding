class Node<V> {
  V val;
  Node<V>? next;

  Node(this.val, [this.next]);
}

class LinkedList<V> {
  Node<V>? head;

  void append(V val) {
    var newNode = Node<V>(val);
    if (head == null) {
      head = newNode;
      return;
    }
    var current = head;
    while (current!.next != null) {
      current = current.next;
    }
    current.next = newNode;
  }

  void printList() {
    var current = head;
    while (current != null) {
      print('${current.val} -> ');
      current = current.next;
    }
    print('nil');
  }
}

void main() {
  var list = LinkedList<int>();
  list.append(1);
  list.append(2);
  list.append(3);
  list.printList();  // Output: 1 -> 2 -> 3 -> nil
}