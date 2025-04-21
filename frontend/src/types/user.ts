export type User = {
  id: string;
  username: string;
  displayName: string;
  email: string;
  createdAt: Date;
  updatedAt: string;
};

export type UserPreview = {
  username: string;
  displayName: string;
  createdAt: Date;
  description: string;
};

export type UserSession = {
  activeToken: string;
  user: UserPreview;
};
