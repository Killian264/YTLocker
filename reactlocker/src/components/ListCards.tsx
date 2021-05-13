import { Channel, Playlist, Video } from "../shared/types";
import { Card } from "./Card";
import { ChannelListItem } from "./ChannelListItem";
import { PlaylistListItem } from "./PlaylistListItem";
import { PlusButton } from "./PlusButton";
import { VideoListItem } from "./VideosListItem";

interface PlaylistListProps {
	className?: string;
	playlists: Playlist[];
}

export const PlaylistList: React.FC<PlaylistListProps> = ({ className, playlists }) => {
	let list = playlists.map((playlist, index) => {
		if (index >= 5) {
			return "";
		}

		return <PlaylistListItem key={index} playlist={playlist}></PlaylistListItem>;
	});

	return (
		<Card className={className}>
			<div className="flex justify-between -mb-1 -mt-1">
				<div className="text-2xl font-semibold">
					<span className="leading-none -mt-0.5">Playlists</span>
				</div>
				<PlusButton color="secondary" disabled={false}></PlusButton>
			</div>
			<div className="grid gap-2">{list}</div>
		</Card>
	);
};

interface VideoListProps {
	className?: string;
	videos: Video[];
}

export const VideoList: React.FC<VideoListProps> = ({ className, videos }) => {
	let list = videos.map((video, index) => {
		if (index >= 5) {
			return "";
		}

		return <VideoListItem key={index} video={video}></VideoListItem>;
	});

	return (
		<Card className={className}>
			<div className="flex justify-between -mb-1 -mt-1">
				<div className="text-2xl font-semibold">
					<span className="leading-none -mt-0.5">Videos</span>
				</div>
			</div>
			<div className="grid gap-2">{list}</div>
		</Card>
	);
};

interface ChannelsListProp {
	className?: string;
	channels: Channel[];
}

export const ChannelsList: React.FC<ChannelsListProp> = ({ className, channels }) => {
	let list = channels.map((channel, index) => {
		if (index >= 5) {
			return "";
		}

		return <ChannelListItem key={index} channel={channel}></ChannelListItem>;
	});

	return (
		<Card className={className}>
			<div className="flex justify-between -mb-1 -mt-1">
				<div className="text-2xl font-semibold">
					<span className="leading-none -mt-0.5">Channels</span>
				</div>
			</div>
			<div className="grid gap-2">{list}</div>
		</Card>
	);
};
