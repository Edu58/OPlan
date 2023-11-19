defmodule Oplan.EventFixtures do
  @moduledoc """
  This module defines test helpers for creating
  entities via the `Oplan.Event` context.
  """

  @doc """
  Generate a event_rating.
  """
  def event_rating_fixture(attrs \\ %{}) do
    {:ok, event_rating} =
      attrs
      |> Enum.into(%{
        event_ratings: "some event_ratings",
        rating: 42
      })
      |> Oplan.Event.create_event_rating()

    event_rating
  end

  @doc """
  Generate a event_comment.
  """
  def event_comment_fixture(attrs \\ %{}) do
    {:ok, event_comment} =
      attrs
      |> Enum.into(%{
        comment: "some comment"
      })
      |> Oplan.Event.create_event_comment()

    event_comment
  end

  @doc """
  Generate a event_follower.
  """
  def event_follower_fixture(attrs \\ %{}) do
    {:ok, event_follower} =
      attrs
      |> Enum.into(%{})
      |> Oplan.Event.create_event_follower()

    event_follower
  end
end
