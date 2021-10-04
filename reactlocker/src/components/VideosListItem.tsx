import React from "react";
import { Color, Video } from "../shared/types";
import { ColorBadge } from "./ColorBadge";
import { ExternalLink } from "./Svg";

export interface VideoListItemProps {
	className?: string;
	video: Video;
	url: string;
	color: Color;
}

export const VideoListItem: React.FC<VideoListItemProps> = ({ video, className, url, color }) => {
	const css = `${className} hover:bg-primary-600 rounded-md flex justify-between cursor-pointer overflow-hidden`;
	const imageSize = "md:h-20 md:w-32 h-16 w-24";
	const textSize = "sm:text-md text-md";

	const open = () => {
		window.open(url, "_blank");
	};

	const UpdatedStats = (video: Video): [number, string] => {
		const diff = Math.abs(new Date().valueOf() - video.created.valueOf());
		const hours = Math.floor(diff / 36e5);
		const minutes = Math.floor(diff / 6e4);
		const seconds = Math.floor(diff / 1e3);

		if (hours === 0 && minutes === 0) {
			return [seconds, "seconds ago"];
		}

		if (hours === 0) {
			return [minutes, "minutes ago"];
		}

		return [hours, "hours ago"];
	};

	return (
		<div className={css} onClick={open}>
			<div className="flex p-1 overflow-hidden">
				<img
					src={video.thumbnailUrl}
					alt="Thumbnail"
					className={`${imageSize} flex-shrink-0 rounded-lg object-cover`}
				/>
				<div className="pl-3 flex flex-col">
					<span className={`${textSize} font-semibold whitespace-nowrap`}>{video.title}</span>
					<span className="text-sm text-primary-text-200">
						{UpdatedStats(video)[0] + " " + UpdatedStats(video)[1]}
					</span>
					<div>
						<ColorBadge className="mt-2" color={color}></ColorBadge>
					</div>
				</div>
			</div>
			<div className="mr-2 ml-4 my-auto select-none ml-4">
				<ExternalLink size={24}></ExternalLink>
			</div>
		</div>
	);
};
