import { ColorToColorCSS } from "./colors";

export interface User {
	username: string;
	email: string;
	picture: string;
	joined: Date;
}

export interface StatCard {
	header: string;
	count: number;
	measurement: string;
}

export interface Playlist {
	id: number;
	youtubeId: string;
	thumbnailUrl: string;
	title: string;
	description: string;
	channels: number[];
	videos: number[];
	created: Date;
	color: Color;
}

export interface Video {
	id: number;
	youtubeId: string;
	thumbnailUrl: string;
	title: string;
	description: string;
	created: Date;
}

export interface Account {
	id: number;
	username: string;
	email: string;
	picture: string;
	permissionLevel: "view" | "manage";
}

export interface Channel {
	id: number;
	youtubeId: string;
	thumbnailUrl: string;
	title: string;
	description: string;
	videos: number[];
	created: Date;
}

export type Color = keyof typeof ColorToColorCSS;
