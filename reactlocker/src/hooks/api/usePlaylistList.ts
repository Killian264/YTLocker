import axios from "axios";
import { useQuery } from "react-query";
import { Playlist } from "../../shared/types";

export const usePlaylistList = (): [boolean, Playlist[]] => {
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
				accountId: item.AccountID,
				created: new Date(Date.parse(item.CreatedAt)),
				videos: item.Videos,
				channels: item.Channels,
				color: item.Color,
			};

			return playlist;
		});

		return playlists;
	});
};
