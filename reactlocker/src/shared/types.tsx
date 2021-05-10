export interface User {
	username: string;
	email: string;
	joined: string;
}

export interface UserLogin {
	email: string;
	password: string;
}

export interface UserRegister {
	username: string;
	email: string;
	password: string;
	password2: string;
}

export interface StatCard {
	header: string;
	count: number;
	measurement: string;
}
