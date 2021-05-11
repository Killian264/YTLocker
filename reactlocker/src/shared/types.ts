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

export interface Playlist {
	id: number;
	youtube: string;
	thumbnail: string;
	title: string;
	description: string;
	url: string;
}

export interface Video {
	id: number;
	youtube: string;
	thumbnail: string;
	title: string;
	description: string;
	url: string;
}

export interface Channel {
	id: number;
	youtube: string;
	thumbnail: string;
	title: string;
	description: string;
	url: string;
}