import { Channel, Playlist, UserLogin, UserRegister, Video } from "./types";
import { DROPLET_BASE } from "./env";
import { useBearer } from "./hooks/useBearer";

interface ApiResponse {
	success: boolean;
	message: string;
}

interface InternalResponse extends ApiResponse {
	data: any;
}

interface LoginResponse extends ApiResponse {
	bearer: string;
}

export interface PlaylistListItem extends Playlist {
	Channels: Channel[];
	Videos: Video[];
}

export interface PlaylistListResponse extends ApiResponse {
	items: PlaylistListItem[];
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

export const API = {
	Login,
	PlaylistList,
};

const ParsePlaylistListResponse = (res: InternalResponse): PlaylistListResponse => {
	if (!res.success) {
		return {
			...res,
			items: [],
		};
	}

	let playlists: PlaylistListItem[] = res.data.map((playlistJSON: any) => {
		let thumbnail = playlistJSON.Thumbnails[playlistJSON.Thumbnails.length - 1];

		let playlist: Playlist = {
			id: playlistJSON.ID,
			youtube: playlistJSON.YoutubeID,
			thumbnail: thumbnail.URL,
			title: playlistJSON.Title,
			description: playlistJSON.Description,
			url: "https://www.youtube.com/playlist?list=" + playlistJSON.YoutubeID,
			created: new Date(Date.parse(playlistJSON.CreatedAt)),
		};

		return {
			...playlist,
			Videos: ParseVideos(playlistJSON, playlist),
			Channels: ParseChannels(playlistJSON, playlist),
		};
	});

	return {
		...res,
		items: playlists,
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

const ParseLoginResponse = (data: InternalResponse): LoginResponse => {
	return {
		...data,
		bearer: data.success ? data.data.Bearer : "",
	};
};

const ParseApiResponse = (data: any): InternalResponse => {
	return {
		success: data.Status === 200,
		message: data.Message,
		data: data.Data,
	};
};
