import { Channel, Playlist, User, UserLogin, UserRegister, Video } from "./types";
import { DROPLET_BASE } from "./env";

interface ApiResponse {
	success: boolean;
	message: string;
}

interface InternalResponse extends ApiResponse {
	data: any;
}

interface UserLoginResponse extends ApiResponse {
	bearer: string;
}

export interface UserInformationResponse extends ApiResponse {
	user: User | null;
}

export interface PlaylistListResponse extends ApiResponse {
	playlists: Playlist[];
}

const Login = (user: UserLogin) => {
	const url = DROPLET_BASE + "/user/login";

	const options = {
		method: "POST",
		headers: {},
		body: JSON.stringify(user),
	};

	return fetch(url, options)
		.then((res) => res.json())
		.then((data) => ParseApiResponse(data))
		.then((data) => ParseLoginResponse(data));
};

const Register = (user: UserRegister) => {
	const url = DROPLET_BASE + "/user/register";

	const options = {
		method: "POST",
		headers: {},
		body: JSON.stringify(user),
	};

	return fetch(url, options)
		.then((res) => res.json())
		.then((data) => ParseApiResponse(data))
		.then((data) => ParseRegisterResponse(data));
};

const PlaylistList = (bearer: string) => {
	const url = DROPLET_BASE + "/playlist/list";

	const options = {
		method: "GET",
		headers: {
			Authorization: bearer,
		},
	};

	return fetch(url, options)
		.then((res) => res.json())
		.then((data) => ParseApiResponse(data))
		.then((data) => ParsePlaylistListResponse(data));
};

const UserInformation = (bearer: string) => {
	const url = DROPLET_BASE + "/user/information";

	const options = {
		method: "GET",
		headers: {
			Authorization: bearer,
		},
	};

	return fetch(url, options)
		.then((res) => res.json())
		.then((data) => ParseApiResponse(data))
		.then((data) => ParseUserInformationResponse(data));
};

export const API = {
	UserLogin: Login,
	UserRegister: Register,
	PlaylistList,
	UserInformation,
};

const ParsePlaylistListResponse = (res: InternalResponse): PlaylistListResponse => {
	if (!res.success) {
		return {
			...res,
			playlists: [],
		};
	}

	let playlists: Playlist[] = res.data.map((playlistJSON: any) => {
		let thumbnail = playlistJSON.Thumbnails[playlistJSON.Thumbnails.length - 1];

		let playlist: Playlist = {
			id: playlistJSON.ID,
			youtube: playlistJSON.YoutubeID,
			thumbnail: thumbnail.URL,
			title: playlistJSON.Title,
			description: playlistJSON.Description,
			url: "https://www.youtube.com/playlist?list=" + playlistJSON.YoutubeID,
			created: new Date(Date.parse(playlistJSON.CreatedAt)),
			Videos: [],
			Channels: [],
		};

		return {
			...playlist,
			Videos: ParseVideos(playlistJSON, playlist),
			Channels: ParseChannels(playlistJSON, playlist),
		};
	});

	return {
		...res,
		playlists: playlists,
	};
};

const ParseLoginResponse = (data: InternalResponse): UserLoginResponse => {
	if (!data.success) {
		return {
			...data,
			bearer: "",
		};
	}

	return {
		...data,
		bearer: data.success ? data.data.Bearer : "",
	};
};

const ParseRegisterResponse = (data: InternalResponse): ApiResponse => {
	return {
		...data,
	};
};

const ParseUserInformationResponse = (data: InternalResponse): UserInformationResponse => {
	if (!data.success) {
		return {
			...data,
			user: null,
		};
	}

	return {
		...data,
		user: {
			username: data.data.Username,
			email: data.data.Email,
			joined: new Date(Date.parse(data.data.CreatedAt)),
		},
	};
};

const ParseChannels = (json: any, playlist: Playlist): Channel[] => {
	let channels: Channel[] = json.Channels.map((channel: any) => {
		// let thumbnail =
		// 	channelJSON.Thumbnails[
		// 		channelJSON.Thumbnails.length - 1
		// 	];

		let parsed: Channel = {
			id: channel.ID,
			youtube: channel.YoutubeID,
			thumbnail:
				"https://i.ytimg.com/vi/1PBNAoKd-70/hqdefault.jpg?sqp=-oaymwEXCNACELwBSFryq4qpAwkIARUAAIhCGAE=&rs=AOn4CLCFnLzV-VCKC28TFfjTi5cQL7zXiA",
			title: channel.Title,
			description: channel.Description,
			url: `https://www.youtube.com/channel/${channel.YoutubeID}`,
			created: new Date(Date.parse(channel.CreatedAt)),
		};

		return parsed;
	});

	return channels;
};

const ParseVideos = (json: any, playlist: Playlist): Video[] => {
	const videos = json.Videos.map((video: any) => {
		// let thumbnail =
		// 	videoJSON.Thumbnails[videoJSON.Thumbnails.length - 1];

		let parsed = {
			id: video.ID,
			youtube: video.YoutubeID,
			thumbnail:
				"https://i.ytimg.com/vi/1PBNAoKd-70/hqdefault.jpg?sqp=-oaymwEXCNACELwBSFryq4qpAwkIARUAAIhCGAE=&rs=AOn4CLCFnLzV-VCKC28TFfjTi5cQL7zXiA",
			title: video.Title,
			description: video.Description,
			created: new Date(Date.parse(video.CreatedAt)),
			url: `https://www.youtube.com/watch?v=${video.YoutubeID}&list=${playlist.youtube}`,
		};

		return parsed;
	});

	return videos;
};

const ParseApiResponse = (data: any): InternalResponse => {
	return {
		success: data.Status === 200,
		message: data.Message,
		data: data.Data,
	};
};
