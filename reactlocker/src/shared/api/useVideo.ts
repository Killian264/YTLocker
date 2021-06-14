import axios from "axios";
import { useQuery } from "react-query";
import { Video } from "../types";

export const useVideo = (id: number, enabled = true): [boolean, Video | null] => {
	const { isSuccess, data } = useQuery(["channel", id], () => fetchVideo(id), {
		staleTime: Infinity,
		enabled: enabled,
	});

	if (!isSuccess || data === undefined) {
		return [true, null];
	}

	return [false, data];
};

const fetchVideo = async (id: number): Promise<Video> => {
	return axios.get(`/video/${id}`).then((response) => {
		let { Data: data } = response.data;

		let thumbnail = data.Thumbnails[data.Thumbnails.length - 1];

		let parsed: Video = {
			id: data.ID,
			youtubeId: data.YoutubeID,
			thumbnailUrl: thumbnail.URL,
			title: data.Title,
			description: data.Description,
			created: new Date(Date.parse(data.CreatedAt)),
		};

		return parsed;
	});
};
