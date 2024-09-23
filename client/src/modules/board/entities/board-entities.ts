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

export interface CreateTaskRequest {
  name: string;
  position: number;
  columnid: number;
  description: string;
}
