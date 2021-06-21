import { useState } from "react";
import { Playlist } from "../shared/types";
import { BuildPlaylistUrl } from "../shared/urls";
import { Button } from "./Button";
import { Card } from "./Card";
import { Modal } from "./Modal";
import { Checkmark, Cog, ExternalLink, LeftArrow, SvgBox, Trash } from "./Svg";

export interface PlaylistViewProps {
	className?: string;
	playlist: Playlist;
	DeleteClick: (id: number) => void;
	BackClick: () => void;
}

export const PlaylistView: React.FC<PlaylistViewProps> = ({
	className,
	DeleteClick,
	BackClick,
	playlist,
}) => {
	const [editing, setEditing] = useState(false);
	const [isOpen, setIsOpen] = useState(false);

	const remove = () => {
		setIsOpen(true);
	};

	const swap = () => {
		setEditing(!editing);
	};

	return (
		<>
			{isOpen && (
				<Modal
					confirmMessage={"Yes, I am sure"}
					header={"Are you sure?"}
					body={
						"Playlist information will be deleted, but the playlist will remain accessible on youtube."
					}
					AcceptClick={() => {
						DeleteClick(playlist.id);
						BackClick();
					}}
					RejectClick={() => {
						setIsOpen(false);
					}}
				/>
			)}
			<Card className={`${className} flex flex-col`}>
				<div className="flex justify-between -mb-1 -mt-1">
					<div className="text-2xl font-semibold mt-auto">
						<span className="leading-none -mt-0.5">{playlist.title}</span>
					</div>
					<div></div>
					{editing ? (
						<SvgBox className={`border-green-400 p-0.5`} onClick={swap}>
							<Checkmark className={`text-green-400`} size={26}></Checkmark>
						</SvgBox>
					) : (
						<SvgBox className={`border-primary-200 p-0.5`} onClick={swap}>
							<Cog className="text-primary-200" size={26}></Cog>
						</SvgBox>
					)}
				</div>
				<a href={BuildPlaylistUrl(playlist.youtubeId)} target="_blank" rel="noreferrer">
					<div className={`col-span-6 rounded-lg object-cover w-full bg-black h-40`} />
				</a>
				<div className="flex gap-2 mt-3">
					<div className="flex-grow">
						<div className="md:flex-row flex-col flex gap-2 justify-between">
							<span className="md:text-3xl text-2xl font-semibold block">{playlist.title}</span>
							<div className="gap-2 flex">
								<a
									href={BuildPlaylistUrl(playlist.youtubeId)}
									target="_blank"
									rel="noreferrer"
								>
									<SvgBox>
										<ExternalLink size={26}></ExternalLink>
									</SvgBox>
								</a>
								{editing && (
									<SvgBox className="text-red-500 border-red-500" onClick={remove}>
										<Trash className="text-red-500" size={26}></Trash>
									</SvgBox>
								)}
							</div>
						</div>
						<div style={{ minHeight: "50px" }}>{playlist.description}</div>
					</div>
				</div>
				<div className="flex justify-between mt-auto pt-4">
					<Button size="medium" color="secondary" onClick={BackClick}>
						Back
					</Button>
				</div>
			</Card>
		</>
	);
};
