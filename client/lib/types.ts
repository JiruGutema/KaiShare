
export interface User {
  name: string;
  username: string;
  id: string;
  email: string;
  createdAt: string;
  updatedAt: string;
}

export interface Session {
  user: User
}
