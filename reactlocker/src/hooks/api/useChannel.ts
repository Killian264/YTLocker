import axios from "axios";
import { useQuery } from "react-query";
import { Channel } from "../../shared/types";

export const useChannel = (id: number): [boolean, Channel | null] => {
	const { isLoading, isError, data } = useQuery(["channel", id], () => fetchChannel(id), {
		staleTime: Infinity,
	});

	if (isLoading || isError || data === undefined) {
		return [true, null];
	}

	if (data === null) {
		return [true, null];
	}

	return [false, data];
};

const fetchChannel = async (id: number): Promise<Channel> => {
	return axios.get(`/channel/${id}`).then((response) => {
		let { Data: data } = response.data;

		let thumbnail = data.Thumbnails[data.Thumbnails.length - 1];

		let parsed: Channel = {
			id: data.ID,
			youtubeId: data.YoutubeID,
			thumbnailUrl: thumbnail.URL,
			title: data.Title,
			description: data.Description,
			created: new Date(Date.parse(data.CreatedAt)),
			videos: data.Videos,
		};

		return parsed;
	});
};
