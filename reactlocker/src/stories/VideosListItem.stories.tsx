import { Story } from "@storybook/react";
import { Card } from "../components/Card";
import { VideoListItem } from "../components/VideosListItem";
import { PlusButton } from "../components/PlusButton";
import { Video } from "../shared/types";

export default {
	title: "VideoListItem",
	component: VideoListItem,
};

const videos: Video[] = [
	{
		id: 932423423,
		youtubeId: "PLamdXAekZPYiqLDNQXQTbm4N_cPBmLPyr",
		thumbnailUrl:
			"https://i.ytimg.com/vi/1PBNAoKd-70/hqdefault.jpg?sqp=-oaymwEXCNACELwBSFryq4qpAwkIARUAAIhCGAE=&rs=AOn4CLCFnLzV-VCKC28TFfjTi5cQL7zXiA",
		title: "DogeLog",
		description: "Videos showing Ben Awad as he builds dogehouse.",
		created: new Date(),
	},
];

const Mocked: Story<{}> = ({ ...props }) => {
	let url = "https://www.youtube.com/playlist?list=PLN3n1USn4xlkZgqq9SdgUXPmgpoxUM9QK";

	return (
		<Card>
			<div></div>
			<div>
				<VideoListItem url={url} video={videos[0]}></VideoListItem>
				<VideoListItem url={url} video={videos[0]} className="mt-3"></VideoListItem>
				<VideoListItem url={url} video={videos[0]} className="mt-3"></VideoListItem>
			</div>
		</Card>
	);
};

export const Primary = Mocked.bind({});

Primary.argTypes = {};
