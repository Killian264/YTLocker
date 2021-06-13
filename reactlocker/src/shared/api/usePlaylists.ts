import axios from "axios";
import { useQuery } from "react-query";
import { ColorArray } from "../colors";
import { Playlist } from "../types";

export const usePlaylists = (): [boolean, Playlist[]] => {
	const { isSuccess, data } = useQuery(["playlists"], () => fetchPlaylist());

	if (!isSuccess || data === undefined) {
		return [true, []];
	}

	return [false, data];
};

export const fetchPlaylist = async (): Promise<Playlist[]> => {
	return axios.get("/playlist/list").then((response) => {
		let { Data: data } = response.data;

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
	});
};
