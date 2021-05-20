import React from "react";
import { useQuery } from "react-query";
import { VideoListItem } from "../components/VideosListItem";
import { useBearer } from "../shared/hooks/useBearer";
import { fetchVideo } from "../shared/api/video";
import { BuildVideoUrl } from "../shared/urls";

export interface VideoListItemControllerProps {
	className?: string;
	videoId: number;
}

export const VideoListItemController: React.FC<VideoListItemControllerProps> = ({ className, videoId }) => {
	const [bearer] = useBearer("");

	const { isLoading, isError, data } = useQuery(["video", videoId], () => fetchVideo(bearer, videoId));

	if (isLoading) {
		return <div>Loading...</div>;
	}

	if (isError || data === undefined || data.video == null) {
		return <div>Error...</div>;
	}

	return <VideoListItem video={data.video} url={BuildVideoUrl(data.video.youtubeId)}></VideoListItem>;
};
