interface BoardTask {
  id: number;
  title: string;
  description: string;
  position: number;
  due_date: string;
  created_at: string;
  updated_at: string;
}

interface BoardColumn {
  id: number;
  name: string;
  position: number;
  created_at: string;
  updated_at: string;
  tasks: BoardTask[];
}

export interface BoardEntity {
  id: number;
  name: string;
  created_at: string;
  updated_at: string;
  columns: BoardColumn[];
}

export interface BoardUpdateEntity {
  id: number;
  title: string;
  cards: BoardCardUpdate[];
}

export interface BoardCardUpdate {
  id: number;
  title: string;
  description: string;
}

export interface CreateTaskRequest {
  name: string;
  position: number;
  columnid: number;
  description: string;
}
