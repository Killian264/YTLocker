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

const AuthHeader = (): Headers => {
	const [bearer, setBearer] = useBearer("");

	return new Headers({
		Authorization: bearer,
	});
};

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

const PlaylistList = () => {
	const url = DROPLET_BASE + "/playlist/list";

	const options = {
		method: "GET",
		headers: AuthHeader(),
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

const ParsePlaylistListResponse = (
	playlistList: InternalResponse
): PlaylistListItem[] => {
	let playlists: PlaylistListItem[] = playlistList.data.map(
		(playlistJSON: any) => {
			let thumbnail =
				playlistJSON.Thumbnails[playlistJSON.Thumbnails.length - 1];

			let playlist: Playlist = {
				id: playlistJSON.ID,
				youtube: playlistJSON.YoutubeID,
				thumbnail: thumbnail.URL,
				title: playlistJSON.Title,
				description: playlistJSON.Description,
				url:
					"https://www.youtube.com/playlist?list=" +
					playlistJSON.YoutubeID,
				created: new Date(Date.parse(playlistJSON.CreatedAt)),
			};

			let videos: Video[] = playlistJSON.Videos.map((videoJSON: any) => {
				// let thumbnail =
				// 	videoJSON.Thumbnails[videoJSON.Thumbnails.length - 1];

				let parsed = {
					id: videoJSON.ID,
					youtube: videoJSON.YoutubeID,
					thumbnail: thumbnail.URL,
					title: videoJSON.Title,
					description: videoJSON.Description,
					created: new Date(Date.parse(videoJSON.CreatedAt)),
				};

				return {
					...parsed,
					url:
						"https://www.youtube.com/playlist?list=PLN3n1USn4xlkZgqq9SdgUXPmgpoxUM9QK",
					// url: `https://www.youtube.com/watch?v=${parsed.youtube}&list=${playlist.youtube}`,
				};
			});

			let channels: Channel[] = playlistJSON.Channels.map(
				(channelJSON: any) => {
					// let thumbnail =
					// 	channelJSON.Thumbnails[
					// 		channelJSON.Thumbnails.length - 1
					// 	];

					let parsed = {
						id: channelJSON.ID,
						youtube: channelJSON.YoutubeID,
						thumbnail: thumbnail.URL,
						title: channelJSON.Title,
						description: channelJSON.Description,
						created: new Date(Date.parse(playlistJSON.CreatedAt)),
					};

					return {
						...parsed,
						url:
							"https://www.youtube.com/playlist?list=PLN3n1USn4xlkZgqq9SdgUXPmgpoxUM9QK",
						// url: `https://www.youtube.com/channel/${parsed.youtube}`,
					};
				}
			);

			return {
				...playlist,
				Videos: videos,
				Channel: channels,
			};
		}
	);

	return playlists;
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
