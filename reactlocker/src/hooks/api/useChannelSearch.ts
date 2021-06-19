import axios from "axios";
import { useQuery } from "react-query";
import { Channel } from "../../shared/types";
import { IsValidYTChannelUrl, ParseYTChannelUrl } from "../../shared/urls";

export const useChannelSearch = (url: string): [boolean, Channel | null] => {
	let enabled = false;
	let kind = "";
	let id = "";

	if (IsValidYTChannelUrl(url)) {
		[kind, id] = ParseYTChannelUrl(url);
		enabled = true;
	}

	const { isLoading, isError, data } = useQuery(["channelSearch", kind, id], () => queryChannel(kind, id), {
		staleTime: Infinity,
		enabled: enabled,
	});

	if (isLoading || isError || data === undefined) {
		return [true, null];
	}

	return [false, data];
};

const queryChannel = async (kind: string, id: string): Promise<Channel> => {
	return axios.get(`/channel/search`, { params: { kind: kind, query: id } }).then((response) => {
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
