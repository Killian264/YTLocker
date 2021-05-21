import { useQuery } from "react-query";
import { useHistory } from "react-router-dom";
import { ColorArray } from "../colors";
import { DROPLET_BASE } from "../env";
import { useBearer } from "../hooks/useBearer";
import { Playlist } from "../types";
import { ApiResponse2 } from "./api";

export interface PlaylistListResponse extends ApiResponse2 {
	playlists: Playlist[];
}

export const usePlaylists = (): [boolean, Playlist[]] => {
	const [bearer] = useBearer("");
	let history = useHistory();

	const { isLoading, isError, data } = useQuery(["playlists"], () => fetchPlaylist(bearer));

	if (isLoading || isError || data === undefined) {
		return [true, []];
	}

	if (data.status === 401) {
		history.push("/login");
	}

	return [false, data.playlists];
};

export const fetchPlaylist = async (bearer: string): Promise<PlaylistListResponse> => {
	const url = DROPLET_BASE + "/playlist/list";
	const options = {
		method: "GET",
		headers: {
			Authorization: bearer,
		},
	};

	const res = await (await fetch(url, options)).json();

	let video: PlaylistListResponse = {
		status: res.Status,
		message: res.Message,
		playlists: [],
	};

	if (video.status === 200) {
		video.playlists = parsePlaylistList(res.Data);
	}

	return video;
};

const parsePlaylistList = (data: any): Playlist[] => {
	let playlists: Playlist[] = data.map((item: any, index: number) => {
		let thumbnail = item.Thumbnails[item.Thumbnails.length - 1];

		let playlist: Playlist = {
			id: item.ID,
			youtubeId: item.YoutubeID,
			thumbnailUrl: thumbnail.URL,
			title: item.Title,
			description: item.Description,
			created: new Date(Date.parse(item.CreatedAt)),
			videos: item.Videos,
			channels: item.Channels,
			color: ColorArray[index],
		};

		return playlist;
	});

	return playlists;
};
